// DO NOT EDIT! Code generated.
package listenconfig

import (
	"context"
	"fmt"
	"net"

	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/components/stream"
	"github.com/wzshiming/pipe/internal/ctxcache"
	"github.com/wzshiming/pipe/internal/logger"
)

func init() {
	register.Register("ref", NewListenConfigRefWithConfig)
	register.Register("def", NewListenConfigDefWithConfig)
	register.Register("none", newListenConfigNone)
}

type Config struct {
	Name string
	Def  stream.ListenConfig `json:",omitempty"`
}

func NewListenConfigRefWithConfig(ctx context.Context, conf *Config) stream.ListenConfig {
	o := &ListenConfig{
		Name: conf.Name,
		Def:  conf.Def,
		Ctx:  ctx,
	}
	return o
}

func NewListenConfigDefWithConfig(ctx context.Context, conf *Config) stream.ListenConfig {
	return ListenConfigPut(ctx, conf.Name, conf.Def)
}

func ListenConfigPut(ctx context.Context, name string, def stream.ListenConfig) stream.ListenConfig {
	if def == nil {
		def = ListenConfigNone
	}

	m, ok := ctxcache.GetCacheWithContext(ctx)
	if !ok {
		return ListenConfigNone
	}
	store, _ := m.LoadOrStore("stream.ListenConfig", map[string]stream.ListenConfig{})
	store.(map[string]stream.ListenConfig)[name] = def
	return def
}

func ListenConfigGet(ctx context.Context, name string, defaults stream.ListenConfig) stream.ListenConfig {
	m, ok := ctxcache.GetCacheWithContext(ctx)
	if ok {
		store, _ := m.LoadOrStore("stream.ListenConfig", map[string]stream.ListenConfig{})
		o, ok := store.(map[string]stream.ListenConfig)[name]
		if ok {
			return o
		}
	}

	if defaults != nil {
		return defaults
	}
	return ListenConfigNone
}

var ListenConfigNone _ListenConfigNone

type _ListenConfigNone struct{}

func newListenConfigNone() stream.ListenConfig {
	return ListenConfigNone
}

func (_ListenConfigNone) ListenStream(_ context.Context) (_ net.Listener, error error) {
	logger.Warn("this is none of stream.ListenConfig")

	error = fmt.Errorf("error none")

	return
}

type ListenConfig struct {
	Name string
	Def  stream.ListenConfig
	Ctx  context.Context
}

func (o *ListenConfig) ListenStream(context context.Context) (net.Listener, error) {
	return ListenConfigGet(o.Ctx, o.Name, o.Def).ListenStream(context)
}
