package manager

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)

var (
	ErrNotFunction             = fmt.Errorf("not a function")
	ErrReturnNoParameters      = fmt.Errorf("returns no parameters")
	ErrTooManyReturnParameters = fmt.Errorf("too many return parameters")
	ErrSecondReturnParameters  = fmt.Errorf("the second return parameter must be error")
)
var stdManager = NewDecoderManager()

func Register(kind string, fun interface{}) error {
	return stdManager.Register(kind, fun)
}

func Get(kind string, out0Type reflect.Type) (reflect.Value, bool) {
	return stdManager.Get(kind, out0Type)
}

func LookType(kind string, out0Type reflect.Type) (string, reflect.Type) {
	return stdManager.LookType(kind, out0Type)
}

func HasType(out0Type reflect.Type) bool {
	return stdManager.HasType(out0Type)
}

type decoderManager struct {
	typeName   map[string]reflect.Type
	decoder    map[reflect.Type]map[string]reflect.Value
	pairs      map[string][]*pair
	interfaces map[reflect.Type]struct{}
}

func NewDecoderManager() *decoderManager {
	return &decoderManager{
		typeName:   map[string]reflect.Type{},
		decoder:    map[reflect.Type]map[string]reflect.Value{},
		pairs:      map[string][]*pair{},
		interfaces: map[reflect.Type]struct{}{},
	}
}

func (h *decoderManager) Register(kind string, v interface{}) error {
	fun := reflect.ValueOf(v)
	typ, err := checkFunc(fun)
	if err != nil {
		log.Printf("[ERROR] Register config: %s.%s: %s: %s", typ.PkgPath(), typ.Name(), kind, err)
		return err
	}

	for {
		h.register(kind, typ, fun)
		if typ.Kind() != reflect.Ptr {
			break
		}
		typ = typ.Elem()
	}
	log.Printf("[INFO] Register config: %s.%s: %s", typ.PkgPath(), typ.Name(), kind)
	return nil
}

func (h *decoderManager) register(kind string, typ reflect.Type, fun reflect.Value) {
	typName := strings.Join([]string{typ.PkgPath(), typ.Name()}, ".")
	h.typeName[typName] = typ

	_, ok := h.decoder[typ]
	if !ok {
		h.decoder[typ] = map[string]reflect.Value{}
	}
	h.decoder[typ][kind] = fun

	if typ.Kind() == reflect.Interface {
		h.pairs[kind] = append(h.pairs[kind], &pair{
			out0Type: typ,
			funValue: fun,
		})

		h.interfaces[typ] = struct{}{}
	}
}

func (h *decoderManager) LookType(kind string, out0Type reflect.Type) (string, reflect.Type) {
	if s := strings.SplitN(kind, "@", 2); len(s) == 2 {
		t, ok := h.typeName[s[0]]
		if !ok {
			return kind, out0Type
		}
		kind = s[1]
		return s[1], t
	}
	return kind, out0Type
}

func (h *decoderManager) HasType(out0Type reflect.Type) bool {
	_, ok := h.decoder[out0Type]
	if ok {
		return true
	}

	for i := range h.interfaces {
		if out0Type.AssignableTo(i) {
			return true
		}
	}

	return false
}

func (h *decoderManager) Get(kind string, out0Type reflect.Type) (reflect.Value, bool) {
	m, ok := h.decoder[out0Type]
	if ok {
		fun, ok := m[kind]
		if ok {
			return fun, ok
		}
	}

	pairs, ok := h.pairs[kind]
	if ok {
		for _, pair := range pairs {
			if out0Type.AssignableTo(pair.out0Type) {
				return pair.funValue, true
			}
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
