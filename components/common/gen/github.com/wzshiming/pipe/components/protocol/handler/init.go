// DO NOT EDIT! Code generated.
package reference

import (
	"context"

	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/components/protocol"
	"github.com/wzshiming/pipe/internal/logger"
)

func init() {
	register.Register("ref", NewHandlerRefWithConfig)
	register.Register("def", NewHandlerDefWithConfig)
	register.Register("none", NewHandlerNone)
}

type Config struct {
	Name string
	Def  protocol.Handler `json:",omitempty"`
}

func NewHandlerRefWithConfig(conf *Config) (protocol.Handler, error) {
	o := &Handler{
		Name: conf.Name,
		Def:  conf.Def,
	}
	return o, nil
}

func NewHandlerDefWithConfig(conf *Config) (protocol.Handler, error) {
	HandlerStore[conf.Name] = conf.Def
	return conf.Def, nil
}

var HandlerStore = map[string]protocol.Handler{}

func HandlerFind(name string, defaults protocol.Handler) protocol.Handler {
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

func NewHandlerNone() protocol.Handler {
	return HandlerNone{}
}

func (HandlerNone) ServeProtocol(_ context.Context, _ protocol.Protocol) {
	logger.Warn("this is none of protocol.Handler")
	return
}

type Handler struct {
	Name string
	Def  protocol.Handler
}

func (o *Handler) ServeProtocol(context context.Context, protocol protocol.Protocol) {
	HandlerFind(o.Name, o.Def).ServeProtocol(context, protocol)
}
