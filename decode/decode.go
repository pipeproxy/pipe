package decode

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

var stdDecoder = &decoder{
	decoderManager: stdManager,
}

func Decode(ctx context.Context, config []byte, i interface{}) error {
	return stdDecoder.Decode(ctx, config, i)
}

type decoder struct {
	decoderManager *decoderManager
}

func (d *decoder) Decode(ctx context.Context, config []byte, i interface{}) error {
	v := reflect.ValueOf(i)
	return d.decode(ctx, config, v)
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

func (d *decoder) decodeOther(ctx context.Context, config []byte, v reflect.Value) error {

	switch config[0] {
	case '[':
		if v := d.indirect(v); v.Kind() == reflect.Slice {
			tmp := []json.RawMessage{}
			err := json.Unmarshal(config, &tmp)
			if err != nil {
				return err
			}
			slice := reflect.MakeSlice(v.Type(), 0, len(tmp))
			typ := v.Type().Elem()
			for i := 0; i != len(tmp); i++ {
				c := tmp[i]
				n := reflect.New(typ)
				err = d.decode(ctx, c, n)
				if err != nil {
					return err
				}
				slice = reflect.Append(slice, n.Elem())
			}
			v.Set(slice)
			return nil
		}
	case '{':
		if v := d.indirect(v); v.Kind() == reflect.Struct {
			tmp := map[string]json.RawMessage{}
			err := json.Unmarshal(config, &tmp)
			if err != nil {
				return err
			}
			typ := v.Type()

			n := reflect.New(typ)

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
					vs := reflect.New(f.Type)
					err = d.decode(ctx, c, vs)
					if err != nil {
						return err
					}
					n.Elem().Field(i).Set(vs.Elem())
				}
			}
			v.Set(n.Elem())
			return nil
		}
	}

	err := json.Unmarshal(config, v.Interface())
	if err != nil {
		return err
	}

	return nil
}

func (d *decoder) decode(ctx context.Context, config []byte, v reflect.Value) error {
	if v.Kind() != reflect.Ptr {
		return ErrParsedParameter
	}

	elem := v.Elem()
	if !elem.CanSet() {
		return ErrMustBeAssignable
	}

	config = bytes.TrimSpace(config)

	fun, ok := d.decoderManager.Get(v.Type().Elem())
	if !ok {
		return d.decodeOther(ctx, config, v)
	}

	var nameField struct {
		Name string `json:"@name"`
	}
	err := json.Unmarshal(config, &nameField)
	if err != nil {
		return fmt.Errorf("pipe.decode error: %w", err)
	}
	if nameField.Name == "" {
		return d.decodeOther(ctx, config, v)
	}

	inj := inject.NewInjector(nil)
	inj.Map(reflect.ValueOf(ctx))
	inj.Map(reflect.ValueOf(nameField.Name))
	inj.Map(reflect.ValueOf(config))
	ret, err := inj.Call(fun)
	if err != nil {
		return fmt.Errorf("pipe.decode error: %w", err)
	}

	if len(ret) == 2 {
		errInterface := ret[1].Interface()
		if errInterface != nil {
			err, ok := errInterface.(error)
			if !ok {
				panic("this should not be performed until")
			}
			if err != nil {
				return fmt.Errorf("pipe.decode error: %w", err)
			}
		}
	}

	r := ret[0]
	if r.Kind() == reflect.Interface {
		r = r.Elem()
	}

	v = d.indirectTo(v, r.Type().Elem())

	if v.Kind() != reflect.Interface {
		r = r.Elem()
	}
	v.Set(r)
	return nil
}
