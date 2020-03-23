package reference

import (
	"context"
	"fmt"
	"reflect"
	"sync"

	"github.com/wzshiming/funcfg/define"
)

var (
	mutex sync.RWMutex
	defs  = map[reflect.Type]map[string]reflect.Value{}
)

func Def(ctx context.Context, name string, def define.Self, i interface{}) error {
	d := reflect.ValueOf(def)
	s := reflect.ValueOf(i)
	s.Set(d)

	t := s.Type().Elem()

	mutex.Lock()
	defer mutex.Unlock()

	m, ok := defs[t]
	if !ok {
		m = map[string]reflect.Value{}
		defs[t] = m
	}
	m[name] = d
	return nil
}

func Ref(ctx context.Context, name string, def define.Self, i interface{}) error {

	s := reflect.ValueOf(i)
	t := s.Type().Elem()

	mutex.RLock()
	defer mutex.RUnlock()

	m, ok := defs[t]
	if !ok || !m[name].IsValid() {
		if def == nil {
			return fmt.Errorf("not define %q", name)
		}
		d := reflect.ValueOf(def)
		s.Set(d)
		return nil
	}

	d := reflect.ValueOf(m[name])
	s.Set(d)
	return nil
}
