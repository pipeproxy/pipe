// DO NOT EDIT! Code generated.
package reference

import (
	"context"
	"net"
	"sync"

	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/components/stream"
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

func NewHandlerRefWithConfig(conf *Config) stream.Handler {
	o := &Handler{
		Name: conf.Name,
		Def:  conf.Def,
	}
	return o
}

func NewHandlerDefWithConfig(conf *Config) stream.Handler {
	return HandlerPut(conf.Name, conf.Def)
}

var (
	mut           sync.RWMutex
	_HandlerStore = map[string]stream.Handler{}
)

func HandlerPut(name string, def stream.Handler) stream.Handler {
	if def == nil {
		def = HandlerNone
	}
	mut.Lock()
	_HandlerStore[name] = def
	mut.Unlock()
	return def
}

func HandlerGet(name string, defaults stream.Handler) stream.Handler {
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

func newHandlerNone() stream.Handler {
	return HandlerNone
}

func (_HandlerNone) ServeStream(_ context.Context, _ net.Conn) {
	logger.Warn("this is none of stream.Handler")
	return
}

type Handler struct {
	Name string
	Def  stream.Handler
}

func (o *Handler) ServeStream(context context.Context, conn net.Conn) {
	HandlerGet(o.Name, o.Def).ServeStream(context, conn)
}
