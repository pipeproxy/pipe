package reference

import (
	"context"
	"fmt"
	"reflect"
	"sync"

	"github.com/wzshiming/funcfg/define"
)

var (
	ErrDefEmpty  = fmt.Errorf("def must not be empty")
	ErrNotDefine = fmt.Errorf("not define")
	ErrNotUse    = fmt.Errorf("not use reference")
)

type ctxVal struct {
	mutex sync.RWMutex
	defs  map[reflect.Type]map[string]reflect.Value
	dep   map[string][]reflect.Value
}

type ctxKey int

func With(ctx context.Context) context.Context {
	v := &ctxVal{}
	return context.WithValue(ctx, ctxKey(0), v)
}

func Err(ctx context.Context) error {
	v, ok := get(ctx)
	if !ok {
		return nil
	}
	return v.err()
}

func get(ctx context.Context) (*ctxVal, bool) {
	v := ctx.Value(ctxKey(0))
	if v == nil {
		return nil, false
	}

	if v, ok := v.(*ctxVal); ok {
		return v, true
	}
	return nil, false
}

func Def(ctx context.Context, name string, def define.Self, i interface{}) error {
	if def == nil {
		return ErrDefEmpty
	}
	val, ok := get(ctx)
	if !ok {
		return ErrNotUse
	}

	s := reflect.ValueOf(i).Elem()
	d := reflect.ValueOf(def)
	s.Set(d)

	return val.def(name, s, d)
}

func Ref(ctx context.Context, name string, def define.Self, i interface{}) error {
	val, ok := get(ctx)
	if !ok {
		return ErrNotUse
	}

	s := reflect.ValueOf(i).Elem()

	return val.ref(name, s)
}

func (c *ctxVal) def(name string, s, d reflect.Value) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	t := s.Type()
	if c.defs == nil {
		c.defs = map[reflect.Type]map[string]reflect.Value{}
	}
	m, ok := c.defs[t]
	if !ok {
		m = map[string]reflect.Value{}
		c.defs[t] = m
	}
	m[name] = d
	return c.check(name, d)
}

func (c *ctxVal) ref(name string, s reflect.Value) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	t := s.Type()
	if c.defs == nil || c.defs[t] == nil || !c.defs[t][name].IsValid() {
		if c.dep == nil {
			c.dep = map[string][]reflect.Value{}
		}
		c.dep[name] = append(c.dep[name], s)
		return nil
	}
	d := c.defs[t][name]
	s.Set(d)
	return nil
}

func (c *ctxVal) check(name string, d reflect.Value) error {
	if c.dep == nil || len(c.dep[name]) == 0 {
		return nil
	}

	for _, v := range c.dep[name] {
		v.Set(d)
	}
	c.dep[name] = c.dep[name][:0]
	delete(c.dep, name)
	return nil
}

func (c *ctxVal) err() error {
	if len(c.dep) == 0 {
		return nil
	}
	miss := make([]string, 0, len(c.dep))
	for key := range c.dep {
		miss = append(miss, key)
	}
	return fmt.Errorf("%w: %q", ErrNotDefine, miss)
}
