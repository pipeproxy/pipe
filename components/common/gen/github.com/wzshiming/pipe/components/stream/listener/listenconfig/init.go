// DO NOT EDIT! Code generated.
package reference

import (
	"context"
	"net"

	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/components/stream/listener"
	"github.com/wzshiming/pipe/internal/logger"
)

func init() {
	register.Register("ref", NewListenConfigRefWithConfig)
	register.Register("def", NewListenConfigDefWithConfig)
	register.Register("none", NewListenConfigNone)
}

type Config struct {
	Name string
	Def  listener.ListenConfig `json:",omitempty"`
}

func NewListenConfigRefWithConfig(conf *Config) (listener.ListenConfig, error) {
	o := &ListenConfig{
		Name: conf.Name,
		Def:  conf.Def,
	}
	return o, nil
}

func NewListenConfigDefWithConfig(conf *Config) (listener.ListenConfig, error) {
	ListenConfigStore[conf.Name] = conf.Def
	return conf.Def, nil
}

var ListenConfigStore = map[string]listener.ListenConfig{}

func ListenConfigFind(name string, defaults listener.ListenConfig) listener.ListenConfig {
	o, ok := ListenConfigStore[name]
	if ok {
		return o
	}
	if defaults != nil {
		return defaults
	}
	return ListenConfigNone{}
}

type ListenConfigNone struct{}

func NewListenConfigNone() listener.ListenConfig {
	return ListenConfigNone{}
}

func (ListenConfigNone) ListenStream(_ context.Context) (_ net.Listener, _ error) {
	logger.Warn("this is none of listener.ListenConfig")
	return
}

type ListenConfig struct {
	Name string
	Def  listener.ListenConfig
}

func (o *ListenConfig) ListenStream(context context.Context) (net.Listener, error) {
	return ListenConfigFind(o.Name, o.Def).ListenStream(context)
}
