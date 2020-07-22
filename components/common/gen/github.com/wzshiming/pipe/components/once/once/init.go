// DO NOT EDIT! Code generated.
package reference

import (
	"context"
	"sync"

	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/components/once"
	"github.com/wzshiming/pipe/internal/logger"
)

func init() {
	register.Register("ref", NewOnceRefWithConfig)
	register.Register("def", NewOnceDefWithConfig)
	register.Register("none", newOnceNone)
}

type Config struct {
	Name string
	Def  once.Once `json:",omitempty"`
}

func NewOnceRefWithConfig(conf *Config) once.Once {
	o := &Once{
		Name: conf.Name,
		Def:  conf.Def,
	}
	return o
}

func NewOnceDefWithConfig(conf *Config) once.Once {
	return OncePut(conf.Name, conf.Def)
}

var (
	mut        sync.RWMutex
	_OnceStore = map[string]once.Once{}
)

func OncePut(name string, def once.Once) once.Once {
	if def == nil {
		def = OnceNone
	}
	mut.Lock()
	_OnceStore[name] = def
	mut.Unlock()
	return def
}

func OnceGet(name string, defaults once.Once) once.Once {
	mut.RLock()
	o, ok := _OnceStore[name]
	mut.RUnlock()
	if ok {
		return o
	}
	if defaults != nil {
		return defaults
	}
	return OnceNone
}

var OnceNone _OnceNone

type _OnceNone struct{}

func newOnceNone() once.Once {
	return OnceNone
}

func (_OnceNone) Do(_ context.Context) (_ error) {
	logger.Warn("this is none of once.Once")
	return
}

type Once struct {
	Name string
	Def  once.Once
}

func (o *Once) Do(context context.Context) error {
	return OnceGet(o.Name, o.Def).Do(context)
}
