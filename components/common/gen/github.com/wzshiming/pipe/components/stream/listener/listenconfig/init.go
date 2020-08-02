// DO NOT EDIT! Code generated.
package reference

import (
	"context"
	"fmt"
	"net"
	"sync"

	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/components/stream/listener"
	"github.com/wzshiming/pipe/internal/logger"
)

func init() {
	register.Register("ref", NewListenConfigRefWithConfig)
	register.Register("def", NewListenConfigDefWithConfig)
	register.Register("none", newListenConfigNone)
}

type Config struct {
	Name string
	Def  listener.ListenConfig `json:",omitempty"`
}

func NewListenConfigRefWithConfig(conf *Config) listener.ListenConfig {
	o := &ListenConfig{
		Name: conf.Name,
		Def:  conf.Def,
	}
	return o
}

func NewListenConfigDefWithConfig(conf *Config) listener.ListenConfig {
	return ListenConfigPut(conf.Name, conf.Def)
}

var (
	mut                sync.RWMutex
	_ListenConfigStore = map[string]listener.ListenConfig{}
)

func ListenConfigPut(name string, def listener.ListenConfig) listener.ListenConfig {
	if def == nil {
		def = ListenConfigNone
	}
	mut.Lock()
	_ListenConfigStore[name] = def
	mut.Unlock()
	return def
}

func ListenConfigGet(name string, defaults listener.ListenConfig) listener.ListenConfig {
	mut.RLock()
	o, ok := _ListenConfigStore[name]
	mut.RUnlock()
	if ok {
		return o
	}
	if defaults != nil {
		return defaults
	}
	return ListenConfigNone
}

var ListenConfigNone _ListenConfigNone

type _ListenConfigNone struct{}

func newListenConfigNone() listener.ListenConfig {
	return ListenConfigNone
}

func (_ListenConfigNone) ListenStream(_ context.Context) (_ net.Listener, error error) {
	logger.Warn("this is none of listener.ListenConfig")

	error = fmt.Errorf("error none")

	return
}

type ListenConfig struct {
	Name string
	Def  listener.ListenConfig
}

func (o *ListenConfig) ListenStream(context context.Context) (net.Listener, error) {
	return ListenConfigGet(o.Name, o.Def).ListenStream(context)
}
