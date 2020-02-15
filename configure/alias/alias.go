package alias

import (
	"fmt"
	"log"
	"reflect"
)

var stdAlias = newAlias()

func Register(name string, typ interface{}) error {
	return stdAlias.Register(name, typ)
}

func Get(typ interface{}) string {
	return stdAlias.Get(typ)
}

func GetType(t reflect.Type) string {
	return stdAlias.GetType(t)
}

type alias struct {
	names map[reflect.Type]string
}

func newAlias() *alias {
	return &alias{
		names: map[reflect.Type]string{},
	}
}

func (a *alias) Register(name string, typ interface{}) error {
	t := reflect.ValueOf(typ)
	for (t.Kind() == reflect.Interface && t.NumMethod() == 0) || t.Kind() == reflect.Ptr {
		if t.IsNil() {
			log.Printf("[ERROR] Register alias: %s: %v", name, typ)
			return fmt.Errorf("nil interface %s", t)
		}
		t = t.Elem()
	}
	log.Printf("[INFO] Register alias: %s: %s", name, GetDefaultName(t.Type()))
	a.names[t.Type()] = name
	return nil
}

func (a *alias) Get(typ interface{}) string {
	return a.GetValue(reflect.ValueOf(typ))
}

func (a *alias) GetValue(v reflect.Value) string {
	if v.Kind() == reflect.Interface && v.NumMethod() == 0 {
		if !v.IsNil() {
			v = v.Elem()
		}
	}
	return a.GetType(v.Type())
}

func (a *alias) GetType(t reflect.Type) string {
	name, ok := a.names[t]
	if !ok {
		name = GetDefaultName(t)
	}
	return name
}

func GetDefaultName(t reflect.Type) string {
	name := t.PkgPath() + "." + t.Name()
	return name
}
