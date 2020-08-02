// DO NOT EDIT! Code generated.
package reference

import (
	"context"
	"sync"

	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/components/protocol"
	"github.com/wzshiming/pipe/internal/logger"
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

func NewHandlerRefWithConfig(conf *Config) protocol.Handler {
	o := &Handler{
		Name: conf.Name,
		Def:  conf.Def,
	}
	return o
}

func NewHandlerDefWithConfig(conf *Config) protocol.Handler {
	return HandlerPut(conf.Name, conf.Def)
}

var (
	mut           sync.RWMutex
	_HandlerStore = map[string]protocol.Handler{}
)

func HandlerPut(name string, def protocol.Handler) protocol.Handler {
	if def == nil {
		def = HandlerNone
	}
	mut.Lock()
	_HandlerStore[name] = def
	mut.Unlock()
	return def
}

func HandlerGet(name string, defaults protocol.Handler) protocol.Handler {
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
}

func (o *Handler) ServeProtocol(context context.Context, protocol protocol.Protocol) {
	HandlerGet(o.Name, o.Def).ServeProtocol(context, protocol)
}
