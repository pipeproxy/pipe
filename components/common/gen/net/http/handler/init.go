// DO NOT EDIT! Code generated.
package handler

import (
	"context"
	"net/http"

	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/internal/ctxcache"
	"github.com/wzshiming/pipe/internal/logger"
)

func init() {
	register.Register("ref", NewHandlerRefWithConfig)
	register.Register("def", NewHandlerDefWithConfig)
	register.Register("none", newHandlerNone)
}

type Config struct {
	Name string
	Def  http.Handler `json:",omitempty"`
}

func NewHandlerRefWithConfig(ctx context.Context, conf *Config) http.Handler {
	o := &Handler{
		Name: conf.Name,
		Def:  conf.Def,
		Ctx:  ctx,
	}
	return o
}

func NewHandlerDefWithConfig(ctx context.Context, conf *Config) http.Handler {
	return HandlerPut(ctx, conf.Name, conf.Def)
}

func HandlerPut(ctx context.Context, name string, def http.Handler) http.Handler {
	if def == nil {
		return HandlerNone
	}

	m, ok := ctxcache.GetCacheWithContext(ctx)
	if !ok {
		return HandlerNone
	}
	store, _ := m.LoadOrStore("http.Handler", map[string]http.Handler{})
	store.(map[string]http.Handler)[name] = def
	return def
}

func HandlerGet(ctx context.Context, name string, defaults http.Handler) http.Handler {
	m, ok := ctxcache.GetCacheWithContext(ctx)
	if ok {
		store, ok := m.Load("http.Handler")
		if ok {
			o, ok := store.(map[string]http.Handler)[name]
			if ok {
				return o
			}
		}
	}

	if defaults != nil {
		return defaults
	}
	logger.Warnf("http.Handler %q is not defined", name)
	return HandlerNone
}

var HandlerNone _HandlerNone

type _HandlerNone struct{}

func newHandlerNone() http.Handler {
	return HandlerNone
}

func (_HandlerNone) ServeHTTP(_ http.ResponseWriter, _ *http.Request) {
	logger.Warn("this is none of http.Handler")

	return
}

type Handler struct {
	Name string
	Def  http.Handler
	Ctx  context.Context
}

func (o *Handler) ServeHTTP(responsewriter http.ResponseWriter, b *http.Request) {
	HandlerGet(o.Ctx, o.Name, o.Def).ServeHTTP(responsewriter, b)
}
