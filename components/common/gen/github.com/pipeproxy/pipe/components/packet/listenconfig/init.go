// DO NOT EDIT! Code generated.
package listenconfig

import (
	"context"
	"fmt"
	"net"

	"github.com/pipeproxy/pipe/components/common/register"
	"github.com/pipeproxy/pipe/components/packet"
	"github.com/pipeproxy/pipe/internal/ctxcache"
	"github.com/pipeproxy/pipe/internal/logger"
)

func init() {
	register.Register("ref", NewListenConfigRefWithConfig)
	register.Register("def", NewListenConfigDefWithConfig)
	register.Register("none", newListenConfigNone)
}

type Config struct {
	Name string
	Def  packet.ListenConfig `json:",omitempty"`
}

func NewListenConfigRefWithConfig(ctx context.Context, conf *Config) packet.ListenConfig {
	o := &ListenConfig{
		Name: conf.Name,
		Def:  conf.Def,
		Ctx:  ctx,
	}
	return o
}

func NewListenConfigDefWithConfig(ctx context.Context, conf *Config) packet.ListenConfig {
	return ListenConfigPut(ctx, conf.Name, conf.Def)
}

func ListenConfigPut(ctx context.Context, name string, def packet.ListenConfig) packet.ListenConfig {
	if def == nil {
		return ListenConfigNone
	}

	m, ok := ctxcache.GetCacheWithContext(ctx)
	if !ok {
		return ListenConfigNone
	}
	store, _ := m.LoadOrStore("packet.ListenConfig", map[string]packet.ListenConfig{})
	store.(map[string]packet.ListenConfig)[name] = def
	return def
}

func ListenConfigGet(ctx context.Context, name string, defaults packet.ListenConfig) packet.ListenConfig {
	m, ok := ctxcache.GetCacheWithContext(ctx)
	if ok {
		store, ok := m.Load("packet.ListenConfig")
		if ok {
			o, ok := store.(map[string]packet.ListenConfig)[name]
			if ok {
				return o
			}
		}
	}

	if defaults != nil {
		return defaults
	}
	logger.Warnf("packet.ListenConfig %q is not defined", name)
	return ListenConfigNone
}

var ListenConfigNone _ListenConfigNone

type _ListenConfigNone struct{}

func newListenConfigNone() packet.ListenConfig {
	return ListenConfigNone
}

func (_ListenConfigNone) ListenPacket(_ context.Context) (_ net.PacketConn, error error) {
	logger.Warn("this is none of packet.ListenConfig")

	error = fmt.Errorf("error packet.ListenConfig is none")

	return
}

type ListenConfig struct {
	Name string
	Def  packet.ListenConfig
	Ctx  context.Context
}

func (o *ListenConfig) ListenPacket(context context.Context) (net.PacketConn, error) {
	return ListenConfigGet(o.Ctx, o.Name, o.Def).ListenPacket(context)
}
