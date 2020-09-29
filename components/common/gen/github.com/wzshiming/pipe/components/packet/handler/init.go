// DO NOT EDIT! Code generated.
package handler

import (
	"context"
	"net"

	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/components/packet"
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
	Def  packet.Handler `json:",omitempty"`
}

func NewHandlerRefWithConfig(ctx context.Context, conf *Config) packet.Handler {
	o := &Handler{
		Name: conf.Name,
		Def:  conf.Def,
		Ctx:  ctx,
	}
	return o
}

func NewHandlerDefWithConfig(ctx context.Context, conf *Config) packet.Handler {
	return HandlerPut(ctx, conf.Name, conf.Def)
}

func HandlerPut(ctx context.Context, name string, def packet.Handler) packet.Handler {
	if def == nil {
		return HandlerNone
	}

	m, ok := ctxcache.GetCacheWithContext(ctx)
	if !ok {
		return HandlerNone
	}
	store, _ := m.LoadOrStore("packet.Handler", map[string]packet.Handler{})
	store.(map[string]packet.Handler)[name] = def
	return def
}

func HandlerGet(ctx context.Context, name string, defaults packet.Handler) packet.Handler {
	m, ok := ctxcache.GetCacheWithContext(ctx)
	if ok {
		store, ok := m.Load("packet.Handler")
		if ok {
			o, ok := store.(map[string]packet.Handler)[name]
			if ok {
				return o
			}
		}
	}

	if defaults != nil {
		return defaults
	}
	logger.Warnf("packet.Handler %q is not defined", name)
	return HandlerNone
}

var HandlerNone _HandlerNone

type _HandlerNone struct{}

func newHandlerNone() packet.Handler {
	return HandlerNone
}

func (_HandlerNone) ServePacket(_ context.Context, _ net.PacketConn) {
	logger.Warn("this is none of packet.Handler")

	return
}

type Handler struct {
	Name string
	Def  packet.Handler
	Ctx  context.Context
}

func (o *Handler) ServePacket(context context.Context, packetconn net.PacketConn) {
	HandlerGet(o.Ctx, o.Name, o.Def).ServePacket(context, packetconn)
}
