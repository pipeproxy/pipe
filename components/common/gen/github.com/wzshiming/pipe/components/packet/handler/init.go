// DO NOT EDIT! Code generated.
package reference

import (
	"context"
	"net"

	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/components/packet"
	"github.com/wzshiming/pipe/internal/logger"
)

func init() {
	register.Register("ref", NewHandlerRefWithConfig)
	register.Register("def", NewHandlerDefWithConfig)
	register.Register("none", NewHandlerNone)
}

type Config struct {
	Name string
	Def  packet.Handler `json:",omitempty"`
}

func NewHandlerRefWithConfig(conf *Config) (packet.Handler, error) {
	o := &Handler{
		Name: conf.Name,
		Def:  conf.Def,
	}
	return o, nil
}

func NewHandlerDefWithConfig(conf *Config) (packet.Handler, error) {
	HandlerStore[conf.Name] = conf.Def
	return conf.Def, nil
}

var HandlerStore = map[string]packet.Handler{}

func HandlerFind(name string, defaults packet.Handler) packet.Handler {
	o, ok := HandlerStore[name]
	if ok {
		return o
	}
	if defaults != nil {
		return defaults
	}
	return HandlerNone{}
}

type HandlerNone struct{}

func NewHandlerNone() packet.Handler {
	return HandlerNone{}
}

func (HandlerNone) ServePacket(_ context.Context, _ net.PacketConn) {
	logger.Warn("this is none of packet.Handler")
	return
}

type Handler struct {
	Name string
	Def  packet.Handler
}

func (o *Handler) ServePacket(context context.Context, packetconn net.PacketConn) {
	HandlerFind(o.Name, o.Def).ServePacket(context, packetconn)
}
