// DO NOT EDIT! Code generated.
package handler

import (
	"context"

	"github.com/pipeproxy/pipe/components/common/register"
	"github.com/pipeproxy/pipe/components/protocol"
	"github.com/pipeproxy/pipe/internal/ctxcache"
	"github.com/pipeproxy/pipe/internal/logger"
)

func init() {
	register.Register("ref", NewHandlerRefWithConfig)
	register.Register("def", NewHandlerDefWithConfig)
	register.Register("none", newHandlerNone)
}

type Config struct {
	Name string
	Def  protocol.Handler `json:",omitempty"`
}

func NewHandlerRefWithConfig(ctx context.Context, conf *Config) protocol.Handler {
	o := &Handler{
		Name: conf.Name,
		Def:  conf.Def,
		Ctx:  ctx,
	}
	return o
}

func NewHandlerDefWithConfig(ctx context.Context, conf *Config) protocol.Handler {
	return HandlerPut(ctx, conf.Name, conf.Def)
}

func HandlerPut(ctx context.Context, name string, def protocol.Handler) protocol.Handler {
	if def == nil {
		return HandlerNone
	}

	m, ok := ctxcache.GetCacheWithContext(ctx)
	if !ok {
		return HandlerNone
	}
	store, _ := m.LoadOrStore("protocol.Handler", map[string]protocol.Handler{})
	store.(map[string]protocol.Handler)[name] = def
	return def
}

func HandlerGet(ctx context.Context, name string, defaults protocol.Handler) protocol.Handler {
	m, ok := ctxcache.GetCacheWithContext(ctx)
	if ok {
		store, ok := m.Load("protocol.Handler")
		if ok {
			o, ok := store.(map[string]protocol.Handler)[name]
			if ok {
				return o
			}
		}
	}

	if defaults != nil {
		return defaults
	}
	logger.Warnf("protocol.Handler %q is not defined", name)
	return HandlerNone
}

var HandlerNone _HandlerNone

type _HandlerNone struct{}

func newHandlerNone() protocol.Handler {
	return HandlerNone
}

func (_HandlerNone) ServeProtocol(_ context.Context, _ protocol.Protocol) {
	logger.Warn("this is none of protocol.Handler")

	return
}

type Handler struct {
	Name string
	Def  protocol.Handler
	Ctx  context.Context
}

func (o *Handler) ServeProtocol(context context.Context, protocol protocol.Protocol) {
	HandlerGet(o.Ctx, o.Name, o.Def).ServeProtocol(context, protocol)
}
