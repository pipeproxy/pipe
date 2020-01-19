package decode

import (
	"fmt"
	"log"
	"reflect"
)

var (
	ErrNotFunction             = fmt.Errorf("not a function")
	ErrReturnNoParameters      = fmt.Errorf("returns no parameters")
	ErrTooManyReturnParameters = fmt.Errorf("too many return parameters")
	ErrSecondReturnParameters  = fmt.Errorf("the second return parameter must be error")
)
var stdManager = newDecoderManager()

func Register(fun interface{}) {
	stdManager.Register(fun)
}

type decoderManager struct {
	decoder map[reflect.Type]reflect.Value
	pairs   []*pair
}

func newDecoderManager() *decoderManager {
	return &decoderManager{
		decoder: map[reflect.Type]reflect.Value{},
	}
}

func (h *decoderManager) Register(v interface{}) error {
	fun := reflect.ValueOf(v)
	typ, err := checkFunc(fun)
	if err != nil {
		log.Printf("[ERROR] Register config decoder: %s: %s.%s: %s", typ.Kind(), typ.PkgPath(), typ.Name(), err)
		return err
	}

	for {
		h.register(typ, fun)
		log.Printf("[INFO] Register config decoder: %s: %s.%s", typ.Kind(), typ.PkgPath(), typ.Name())
		if typ.Kind() != reflect.Ptr {
			break
		}
		typ = typ.Elem()
	}
	return nil
}

func (h *decoderManager) register(typ reflect.Type, fun reflect.Value) {
	if typ.Kind() == reflect.Interface {
		h.pairs = append(h.pairs, &pair{
			out0Type: typ,
			funValue: fun,
		})
	} else {
		h.decoder[typ] = fun
	}
}

func (h *decoderManager) Get(out0Type reflect.Type) (reflect.Value, bool) {
	fun, ok := h.decoder[out0Type]
	if ok {
		return fun, ok
	}
	for _, pair := range h.pairs {
		if out0Type.AssignableTo(pair.out0Type) {
			return pair.funValue, true
		}
	}
	return reflect.Value{}, false
}

type pair struct {
	out0Type reflect.Type
	funValue reflect.Value
}

func checkFunc(funcValue reflect.Value) (reflect.Type, error) {
	if funcValue.Kind() != reflect.Func {
		return nil, ErrNotFunction
	}
	funcType := funcValue.Type()

	numOut := funcType.NumOut()
	switch numOut {
	case 0:
		return nil, ErrReturnNoParameters
	case 1:
	case 2:
		if !funcType.Out(1).Implements(errImplements) {
			return nil, ErrSecondReturnParameters
		}
	default:
		return nil, ErrTooManyReturnParameters
	}
	out0Type := funcType.Out(0)

	return out0Type, nil
}

var errImplements = reflect.TypeOf(new(error)).Elem()
