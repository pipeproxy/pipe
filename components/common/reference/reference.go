package reference

import (
	"fmt"
	"reflect"
	"sync"

	"github.com/wzshiming/funcfg/define"
)

var (
	mutex sync.RWMutex
	defs  = map[reflect.Type]map[string]reflect.Value{}
)

func Def(name string, def define.Self, i interface{}) error {
	s := reflect.ValueOf(i).Elem()
	d := reflect.ValueOf(def)
	s.Set(d)

	t := s.Type()

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

func Ref(name string, def define.Self, i interface{}) error {
	s := reflect.ValueOf(i).Elem()
	d := reflect.ValueOf(def)

	t := s.Type()

	mutex.Lock()
	defer mutex.Unlock()

	m, ok := defs[t]
	if !ok {
		if def == nil {
			return fmt.Errorf("not define %q", name)
		}
	} else {
		d = m[name]
	}
	s.Set(d)
	return nil
}
