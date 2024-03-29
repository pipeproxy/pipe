// DO NOT EDIT! Code generated.
package once

import (
	"context"
	"fmt"

	"github.com/pipeproxy/pipe/components/common/register"
	"github.com/pipeproxy/pipe/components/once"
	"github.com/pipeproxy/pipe/internal/ctxcache"
	"github.com/wzshiming/logger"
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

func NewOnceRefWithConfig(ctx context.Context, conf *Config) once.Once {
	o := &Once{
		Name: conf.Name,
		Def:  conf.Def,
		Ctx:  ctx,
	}
	return o
}

func NewOnceDefWithConfig(ctx context.Context, conf *Config) once.Once {
	return OncePut(ctx, conf.Name, conf.Def)
}

func OncePut(ctx context.Context, name string, def once.Once) once.Once {
	if def == nil {
		return OnceNone
	}

	m, ok := ctxcache.GetCacheWithContext(ctx)
	if !ok {
		return OnceNone
	}
	store, _ := m.LoadOrStore("once.Once", map[string]once.Once{})
	store.(map[string]once.Once)[name] = def
	return def
}

func OnceGet(ctx context.Context, name string, defaults once.Once) once.Once {
	m, ok := ctxcache.GetCacheWithContext(ctx)
	if ok {
		store, ok := m.Load("once.Once")
		if ok {
			o, ok := store.(map[string]once.Once)[name]
			if ok {
				return o
			}
		}
	}

	if defaults != nil {
		return defaults
	}
	logger.FromContext(ctx).V(-1).Info("once.Once is not defined", "name", name)
	return OnceNone
}

var OnceNone _OnceNone

type _OnceNone struct{}

func newOnceNone() once.Once {
	return OnceNone
}

func (_OnceNone) Do(_ context.Context) (error error) {
	logger.Log.V(-1).Info("this is none of once.Once")

	error = fmt.Errorf("error once.Once is none")

	return
}

type Once struct {
	Name string
	Def  once.Once
	Ctx  context.Context
}

func (o *Once) Do(context context.Context) error {
	return OnceGet(o.Ctx, o.Name, o.Def).Do(context)
}
