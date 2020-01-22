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
		v := d.indirect(v)
		switch v.Kind() {
		case reflect.Slice:
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
		v := d.indirect(v)
		switch v.Kind() {
		case reflect.Map:
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
		case reflect.Struct:
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

	if !d.decoderManager.HasType(v.Type().Elem()) {
		return d.decodeOther(ctx, config, v)
	}

	var field struct {
		Kind string `json:"@Kind"`
	}
	err := json.Unmarshal(config, &field)
	if err != nil || field.Kind == "" {
		return d.decodeOther(ctx, config, v)
	}

	fun, ok := d.decoderManager.Get(field.Kind, v.Type().Elem())
	if !ok {
		return fmt.Errorf("not defined name %q of %s", field.Kind, v.Type().Elem())
	}

	inj := inject.NewInjector(nil)
	args := []interface{}{&ctx, field.Kind, config}
	for _, arg := range args {
		err := inj.Map(reflect.ValueOf(arg))
		if err != nil {
			return fmt.Errorf("pipe.configure error: %w", err)
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
			return fmt.Errorf("config %q error: %w", config, err)
		}
		err = inj.Map(n)
		if err != nil {
			return fmt.Errorf("pipe.configure error: %w", err)
		}
	}

	ret, err := inj.Call(fun)
	if err != nil {
		return fmt.Errorf("pipe.configure error: %w", err)
	}

	if len(ret) == 2 {
		errInterface := ret[1].Interface()
		if errInterface != nil {
			err, ok := errInterface.(error)
			if !ok {
				panic("this should not be performed until")
			}
			if err != nil {
				return fmt.Errorf("pipe.configure error: %w", err)
			}
		}
	}

	r := ret[0]
	if r.Kind() == reflect.Interface {
		r = r.Elem()
	}

	typ := r.Type()
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	v = d.indirectTo(v, typ)
	if v.Kind() != reflect.Interface {
		r = r.Elem()
	}
	v.Set(r)
	return nil
}
