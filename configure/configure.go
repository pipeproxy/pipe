package configure

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/wzshiming/inject"
	"github.com/wzshiming/pipe/configure/manager"
)

var (
	ErrParsedParameter  = fmt.Errorf("the parsed parameter must be a pointer")
	ErrMustBeAssignable = fmt.Errorf("must be assignable")
)

func Decode(ctx context.Context, config []byte, i interface{}) error {
	return newDecoder().Decode(ctx, config, i)
}

type decoder struct {
	refs         map[string]reflect.Value
	exists       map[string]struct{}
	dependentRef map[string][]func() error
}

func newDecoder() *decoder {
	return &decoder{
		refs:         map[string]reflect.Value{},
		exists:       map[string]struct{}{},
		dependentRef: map[string][]func() error{},
	}
}

func (d *decoder) Decode(ctx context.Context, config []byte, i interface{}) error {
	v := reflect.ValueOf(i)
	_, err := d.decode(ctx, config, v)
	if err != nil {
		return err
	}

	need := []string{}
	for name := range d.dependentRef {
		if _, ok := d.exists[name]; !ok {
			need = append(need, name)
		}
	}
	if len(need) != 0 {
		return fmt.Errorf("missing dependency %v", need)
	}
	return nil
}

func (d *decoder) dependent(names []string, todo func() error) error {
	switch l := len(names); l {
	case 0:
		return todo()
	case 1:
		name := names[0]
		d.dependentRef[name] = append(d.dependentRef[name], todo)
	default:
		i := 0
		for _, name := range names {
			d.dependentRef[name] = append(d.dependentRef[name], func() error {
				i++
				if i == l {
					return todo()
				}
				return nil
			})
		}
	}

	return nil
}

func (d *decoder) setExists(name string) {
	d.exists[name] = struct{}{}
}

func (d *decoder) register(name string, v reflect.Value) error {
	if _, ok := d.refs[name]; ok {
		return fmt.Errorf("duplicate name %q", name)
	}
	d.refs[name] = v
	if dep, ok := d.dependentRef[name]; ok {
		for _, d := range dep {
			err := d()
			if err != nil {
				return err
			}
		}
		delete(d.dependentRef, name)
	}
	return nil
}

func (d *decoder) lookAt(name string) (reflect.Value, bool) {
	v, ok := d.refs[name]
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

func (d *decoder) decodeSlice(ctx context.Context, config []byte, v reflect.Value) ([]string, error) {
	tmp := []json.RawMessage{}
	err := json.Unmarshal(config, &tmp)
	if err != nil {
		return nil, err
	}
	slice := reflect.MakeSlice(v.Type(), 0, len(tmp))
	typ := v.Type().Elem()
	deps := []string{}
	for i := 0; i != len(tmp); i++ {
		slice = reflect.Append(slice, reflect.Zero(typ))
		dep, err := d.decode(ctx, tmp[i], slice.Index(i).Addr())
		if err != nil {
			return nil, err
		}
		deps = append(deps, dep...)
	}
	v.Set(slice)
	return deps, nil
}

func (d *decoder) decodeMap(ctx context.Context, config []byte, v reflect.Value) ([]string, error) {
	tmp := map[string]json.RawMessage{}
	err := json.Unmarshal(config, &tmp)
	if err != nil {
		return nil, err
	}
	typ := v.Type()
	n := reflect.MakeMap(typ)
	deps := []string{}
	for key, raw := range tmp {
		val := reflect.New(typ.Elem())
		dep, err := d.decode(ctx, raw, val)
		if err != nil {
			return nil, err
		}
		deps = append(deps, dep...)
		n.SetMapIndex(reflect.ValueOf(key), val.Elem())
	}
	v.Set(n)
	return deps, nil
}

func (d *decoder) decodeStruct(ctx context.Context, config []byte, v reflect.Value) ([]string, error) {
	tmp := map[string]json.RawMessage{}
	err := json.Unmarshal(config, &tmp)
	if err != nil {
		return nil, err
	}
	typ := v.Type()
	v.Set(reflect.Zero(typ))
	num := typ.NumField()
	deps := []string{}
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
			field := v.Field(i)
			field.Set(reflect.Zero(f.Type))
			dep, err := d.decode(ctx, c, field.Addr())
			if err != nil {
				return nil, err
			}
			deps = append(deps, dep...)
		}
	}
	return deps, nil
}

func (d *decoder) decodeOther(ctx context.Context, config []byte, v reflect.Value) ([]string, error) {
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
		return nil, err
	}

	return nil, nil
}

func (d *decoder) ref(name, ref string, value reflect.Value) error {
	v, ok := d.lookAt(ref)
	if ok {
		err := d.set(value, v.Type(), v)
		if err != nil {
			return fmt.Errorf("ref %s: error %w", ref, err)
		}
		if name != "" {
			err := d.register(name, v)
			if err != nil {
				return err
			}
		}
		return nil
	}
	return fmt.Errorf("not defined name %q", ref)
}

func (d *decoder) kind(ctx context.Context, name string, kind string, typ reflect.Type, config []byte, value reflect.Value) ([]string, error) {
	fun, ok := manager.Get(kind, typ)
	if !ok {
		return nil, fmt.Errorf("not defined name %q of %s", kind, value.Type().Elem())
	}

	inj := inject.NewInjector(nil)
	args := []interface{}{&ctx, kind, config}
	for _, arg := range args {
		err := inj.Map(reflect.ValueOf(arg))
		if err != nil {
			return nil, fmt.Errorf("pipe.configure error: %w", err)
		}
	}
	funType := fun.Type()
	num := funType.NumIn()
	deps := []string{}
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
		dep, err := d.decodeOther(ctx, config, n)
		if err != nil {
			return nil, fmt.Errorf("config %q error: %w", config, err)
		}
		err = inj.Map(n)
		if err != nil {
			return nil, fmt.Errorf("pipe.configure map args error: %w", err)
		}
		deps = append(deps, dep...)

	}

	if len(deps) == 0 {
		err := d.call(inj, fun, name, typ, value)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}

	err := d.dependent(deps, func() error {
		return d.call(inj, fun, name, typ, value)
	})
	if err != nil {
		return nil, err
	}
	return deps, nil
}

func (d *decoder) call(inj *inject.Injector, fun reflect.Value, name string, typ reflect.Type, value reflect.Value) error {
	ret, err := inj.Call(fun)
	if err != nil {
		return fmt.Errorf("pipe.configure call error: %w", err)
	}

	if len(ret) == 2 {
		errInterface := ret[1].Interface()
		if errInterface != nil {
			err, ok := errInterface.(error)
			if !ok {
				panic("this should not be performed until")
			}
			if err != nil {
				return fmt.Errorf("pipe.configure call error: %w", err)
			}
		}
	}

	r := ret[0]
	if name != "" {
		err := d.register(name, r)
		if err != nil {
			return err
		}
	}

	err = d.set(value, typ, r)
	if err != nil {
		return err
	}
	return nil
}

func (d *decoder) decode(ctx context.Context, config []byte, value reflect.Value) ([]string, error) {
	if value.Kind() != reflect.Ptr {
		return nil, ErrParsedParameter
	}

	elem := value.Elem()
	if !elem.CanSet() {
		return nil, ErrMustBeAssignable
	}

	var field struct {
		Kind string `json:"@Kind"`

		Name string `json:"@Name"`
		Ref  string `json:"@Ref"`
	}
	err := json.Unmarshal(config, &field)
	if err != nil {
		return d.decodeOther(ctx, config, value)
	}

	if field.Name != "" {
		d.setExists(field.Name)
	}

	if field.Kind == "" && field.Ref == "" {
		return d.decodeOther(ctx, config, value)
	}

	if field.Ref != "" {
		err := d.ref(field.Name, field.Ref, value)
		if err != nil {
			dep := []string{field.Ref}
			err := d.dependent(dep, func() error {
				return d.ref(field.Name, field.Ref, value)
			})
			if err != nil {
				return nil, err
			}
			return dep, nil
		}
		return nil, nil
	}

	kind := field.Kind
	typ := value.Type().Elem()
	kind, typ = manager.LookType(kind, typ)

	if kind == "" {
		return d.decodeOther(ctx, config, value)
	}

	if !manager.HasType(typ) {
		return nil, fmt.Errorf("not define config %s %v", kind, typ)
	}

	deps, err := d.kind(ctx, field.Name, kind, typ, config, value)
	if err != nil {
		return nil, fmt.Errorf("config %q: error %w", config, err)
	}

	return deps, nil
}

func (d *decoder) set(value reflect.Value, typ reflect.Type, r reflect.Value) error {
	if r.Kind() == reflect.Interface {
		r = r.Elem()
	}

	if r.Kind() == reflect.Invalid {
		return fmt.Errorf("got %s, want %s", r.String(), value.String())
	}

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	value = d.indirectTo(value, typ)
	if value.Kind() != reflect.Interface {
		r = r.Elem()
	}

	switch value.Kind() {
	case reflect.Interface:
		if !r.Type().Implements(value.Type()) {
			return fmt.Errorf("value of %s is not assignable to %s", r.Type(), value.Type())
		}
	default:
		if !r.Type().AssignableTo(value.Type()) {
			return fmt.Errorf("value of %s is not assignable to %s", r.Type(), value.Type())
		}
	}

	value.Set(r)
	return nil
}
