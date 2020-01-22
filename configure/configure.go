package configure

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/wzshiming/inject"
)

var (
	ErrParsedParameter  = fmt.Errorf("the parsed parameter must be a pointer")
	ErrMustBeAssignable = fmt.Errorf("must be assignable")
)

var stdDecoder = newDecoder()

func Decode(ctx context.Context, config []byte, i interface{}) error {
	return stdDecoder.Decode(ctx, config, i)
}

type decoder struct {
	decoderManager *decoderManager
	temp           map[string]reflect.Value
	defers         []func() error
}

func newDecoder() *decoder {
	return &decoder{
		decoderManager: stdManager,
		temp:           map[string]reflect.Value{},
	}
}

func (d *decoder) Decode(ctx context.Context, config []byte, i interface{}) error {
	v := reflect.ValueOf(i)
	err := d.decode(ctx, config, v)
	if err != nil {
		return err
	}
	for _, def := range d.defers {
		err = def()
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *decoder) register(name string, v reflect.Value) {
	d.temp[name] = v
}

func (d *decoder) lookAt(name string) (reflect.Value, bool) {
	v, ok := d.temp[name]
	return v, ok
}

func (d *decoder) indirect(v reflect.Value) reflect.Value {
	if v.Kind() != reflect.Ptr {
		return v
	}
	if v.IsNil() {
		v.Set(reflect.New(v.Type().Elem()))
	}
	return d.indirect(v.Elem())
}

func (d *decoder) indirectTo(v reflect.Value, to reflect.Type) reflect.Value {
	if v.Type().AssignableTo(to) {
		return v
	}
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			v.Set(reflect.New(v.Type().Elem()))
		}
		return d.indirectTo(v.Elem(), to)
	}
	return v
}

func (d *decoder) decodeSlice(ctx context.Context, config []byte, v reflect.Value) error {
	tmp := []json.RawMessage{}
	err := json.Unmarshal(config, &tmp)
	if err != nil {
		return err
	}
	slice := reflect.MakeSlice(v.Type(), 0, len(tmp))
	typ := v.Type().Elem()
	for i := 0; i != len(tmp); i++ {
		slice = reflect.Append(slice, reflect.Zero(typ))
		err = d.decode(ctx, tmp[i], slice.Index(i).Addr())
		if err != nil {
			return err
		}
	}
	v.Set(slice)
	return nil
}

func (d *decoder) decodeMap(ctx context.Context, config []byte, v reflect.Value) error {
	tmp := map[string]json.RawMessage{}
	err := json.Unmarshal(config, &tmp)
	if err != nil {
		return err
	}
	typ := v.Type()
	n := reflect.MakeMap(typ)
	for key, raw := range tmp {
		val := reflect.New(typ.Elem())
		err := d.decode(ctx, raw, val)
		if err != nil {
			return err
		}
		n.SetMapIndex(reflect.ValueOf(key), val.Elem())
	}
	v.Set(n)
	return nil
}

func (d *decoder) decodeStruct(ctx context.Context, config []byte, v reflect.Value) error {
	tmp := map[string]json.RawMessage{}
	err := json.Unmarshal(config, &tmp)
	if err != nil {
		return err
	}
	typ := v.Type()
	v.Set(reflect.Zero(typ))
	num := typ.NumField()
	for i := 0; i != num; i++ {
		f := typ.Field(i)
		name := f.Name
		if value, ok := f.Tag.Lookup("json"); ok {
			n := strings.SplitN(value, ",", 2)
			if n[0] != "" {
				name = n[0]
			}
		}

		if c, ok := tmp[name]; ok {
			v.Field(i).Set(reflect.Zero(f.Type))
			err = d.decode(ctx, c, v.Field(i).Addr())
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (d *decoder) decodeOther(ctx context.Context, config []byte, v reflect.Value) error {
	config = bytes.TrimSpace(config)
	switch config[0] {
	case '[':
		v := d.indirect(v)
		switch v.Kind() {
		case reflect.Slice:
			return d.decodeSlice(ctx, config, v)
		}
	case '{':
		v := d.indirect(v)
		switch v.Kind() {
		case reflect.Map:
			return d.decodeMap(ctx, config, v)
		case reflect.Struct:
			return d.decodeStruct(ctx, config, v)
		}
	}

	err := json.Unmarshal(config, v.Interface())
	if err != nil {
		return err
	}

	return nil
}

func (d *decoder) appendDefer(ref string, value reflect.Value) {
	d.defers = append(d.defers, func() error {
		return d.ref(ref, value)
	})
}

func (d *decoder) ref(ref string, value reflect.Value) error {
	v, ok := d.lookAt(ref)
	if ok {
		d.set(value, v)
		return nil
	}
	return fmt.Errorf("not defined name %q", ref)
}

func (d *decoder) getKind(ctx context.Context, kind string, config []byte, value reflect.Value) (reflect.Value, error) {
	fun, ok := d.decoderManager.Get(kind, value.Type().Elem())
	if !ok {
		return reflect.Value{}, fmt.Errorf("not defined name %q of %s", kind, value.Type().Elem())
	}

	inj := inject.NewInjector(nil)
	args := []interface{}{&ctx, kind, config}
	for _, arg := range args {
		err := inj.Map(reflect.ValueOf(arg))
		if err != nil {
			return reflect.Value{}, fmt.Errorf("pipe.configure error: %w", err)
		}
	}
	funType := fun.Type()
	num := funType.NumIn()
	for i := 0; i != num; i++ {
		in := funType.In(i)
		switch in.Kind() {
		case reflect.Slice:
			if in.Elem().Kind() == reflect.Uint8 {
				continue
			}
		case reflect.Struct, reflect.Map:
		case reflect.Ptr:
			if in.Elem().Kind() != reflect.Struct {
				continue
			}
		case reflect.String:
			continue
		default:
			continue
		}

		n := reflect.New(in)
		err := d.decode(ctx, config, n)
		if err != nil {
			return reflect.Value{}, fmt.Errorf("config %q error: %w", config, err)
		}
		err = inj.Map(n)
		if err != nil {
			return reflect.Value{}, fmt.Errorf("pipe.configure error: %w", err)
		}
	}

	ret, err := inj.Call(fun)
	if err != nil {
		return reflect.Value{}, fmt.Errorf("pipe.configure error: %w", err)
	}

	if len(ret) == 2 {
		errInterface := ret[1].Interface()
		if errInterface != nil {
			err, ok := errInterface.(error)
			if !ok {
				panic("this should not be performed until")
			}
			if err != nil {
				return reflect.Value{}, fmt.Errorf("pipe.configure error: %w", err)
			}
		}
	}
	return ret[0], nil
}
func (d *decoder) decode(ctx context.Context, config []byte, value reflect.Value) error {
	if value.Kind() != reflect.Ptr {
		return ErrParsedParameter
	}

	elem := value.Elem()
	if !elem.CanSet() {
		return ErrMustBeAssignable
	}

	if !d.decoderManager.HasType(value.Type().Elem()) {
		return d.decodeOther(ctx, config, value)
	}

	var field struct {
		Kind string `json:"@Kind"`

		Name string `json:"@Name"`
		Ref  string `json:"@Ref"`
	}
	err := json.Unmarshal(config, &field)
	if err != nil || (field.Kind == "" && field.Ref == "") {
		return d.decodeOther(ctx, config, value)
	}

	var r reflect.Value
	if field.Ref != "" {
		err := d.ref(field.Ref, value)
		if err != nil {
			d.appendDefer(field.Ref, value)
		}
		return nil
	}

	if field.Kind != "" && r == (reflect.Value{}) {
		r, err = d.getKind(ctx, field.Kind, config, value)
		if err != nil {
			return err
		}
	}

	if field.Name != "" {
		d.register(field.Name, r)
	}

	d.set(value, r)

	return nil
}

func (d *decoder) set(value, r reflect.Value) {
	if r.Kind() == reflect.Interface {
		r = r.Elem()
	}

	typ := r.Type()
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	value = d.indirectTo(value, typ)
	if value.Kind() != reflect.Interface {
		r = r.Elem()
	}
	value.Set(r)
}
