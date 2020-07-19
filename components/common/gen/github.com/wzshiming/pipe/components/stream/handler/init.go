// DO NOT EDIT! Code generated.
package reference

import (
	"context"
	"net"

	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/components/stream"
	"github.com/wzshiming/pipe/internal/logger"
)

func init() {
	register.Register("ref", NewHandlerRefWithConfig)
	register.Register("def", NewHandlerDefWithConfig)
	register.Register("none", NewHandlerNone)
}

type Config struct {
	Name string
	Def  stream.Handler `json:",omitempty"`
}

func NewHandlerRefWithConfig(conf *Config) (stream.Handler, error) {
	o := &Handler{
		Name: conf.Name,
		Def:  conf.Def,
	}
	return o, nil
}

func NewHandlerDefWithConfig(conf *Config) (stream.Handler, error) {
	HandlerStore[conf.Name] = conf.Def
	return conf.Def, nil
}

var HandlerStore = map[string]stream.Handler{}

func HandlerFind(name string, defaults stream.Handler) stream.Handler {
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

func NewHandlerNone() stream.Handler {
	return HandlerNone{}
}

func (HandlerNone) ServeStream(_ context.Context, _ net.Conn) {
	logger.Warn("this is none of stream.Handler")
	return
}

type Handler struct {
	Name string
	Def  stream.Handler
}

func (o *Handler) ServeStream(context context.Context, conn net.Conn) {
	HandlerFind(o.Name, o.Def).ServeStream(context, conn)
}
