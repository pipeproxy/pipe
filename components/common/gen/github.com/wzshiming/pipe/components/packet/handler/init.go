// DO NOT EDIT! Code generated.
package reference

import (
	"context"
	"net"
	"sync"

	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/components/packet"
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

func NewHandlerRefWithConfig(conf *Config) packet.Handler {
	o := &Handler{
		Name: conf.Name,
		Def:  conf.Def,
	}
	return o
}

func NewHandlerDefWithConfig(conf *Config) packet.Handler {
	return HandlerPut(conf.Name, conf.Def)
}

var (
	mut           sync.RWMutex
	_HandlerStore = map[string]packet.Handler{}
)

func HandlerPut(name string, def packet.Handler) packet.Handler {
	if def == nil {
		def = HandlerNone
	}
	mut.Lock()
	_HandlerStore[name] = def
	mut.Unlock()
	return def
}

func HandlerGet(name string, defaults packet.Handler) packet.Handler {
	mut.RLock()
	o, ok := _HandlerStore[name]
	mut.RUnlock()
	if ok {
		return o
	}
	if defaults != nil {
		return defaults
	}
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
}

func (o *Handler) ServePacket(context context.Context, packetconn net.PacketConn) {
	HandlerGet(o.Name, o.Def).ServePacket(context, packetconn)
}
