// DO NOT EDIT! Code generated.
package handler

import (
	"context"
	"net"

	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/components/stream"
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
	Def  stream.Handler `json:",omitempty"`
}

func NewHandlerRefWithConfig(ctx context.Context, conf *Config) stream.Handler {
	o := &Handler{
		Name: conf.Name,
		Def:  conf.Def,
		Ctx:  ctx,
	}
	return o
}

func NewHandlerDefWithConfig(ctx context.Context, conf *Config) stream.Handler {
	return HandlerPut(ctx, conf.Name, conf.Def)
}

func HandlerPut(ctx context.Context, name string, def stream.Handler) stream.Handler {
	if def == nil {
		return HandlerNone
	}

	m, ok := ctxcache.GetCacheWithContext(ctx)
	if !ok {
		return HandlerNone
	}
	store, _ := m.LoadOrStore("handler.Handler", map[string]stream.Handler{})
	store.(map[string]stream.Handler)[name] = def
	return def
}

func HandlerGet(ctx context.Context, name string, defaults stream.Handler) stream.Handler {
	m, ok := ctxcache.GetCacheWithContext(ctx)
	if ok {
		store, ok := m.Load("handler.Handler")
		if ok {
			o, ok := store.(map[string]stream.Handler)[name]
			if ok {
				return o
			}
		}
	}

	if defaults != nil {
		return defaults
	}
	logger.Warnf("handler.Handler %q is not defined", name)
	return HandlerNone
}

var HandlerNone _HandlerNone

type _HandlerNone struct{}

func newHandlerNone() stream.Handler {
	return HandlerNone
}

func (_HandlerNone) ServeStream(_ context.Context, _ net.Conn) {
	logger.Warn("this is none of handler.Handler")

	return
}

type Handler struct {
	Name string
	Def  stream.Handler
	Ctx  context.Context
}

func (o *Handler) ServeStream(context context.Context, conn net.Conn) {
	HandlerGet(o.Ctx, o.Name, o.Def).ServeStream(context, conn)
}
