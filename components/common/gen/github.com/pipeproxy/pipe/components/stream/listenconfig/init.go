// DO NOT EDIT! Code generated.
package listenconfig

import (
	"context"
	"fmt"
	"net"

	"github.com/pipeproxy/pipe/components/common/register"
	"github.com/pipeproxy/pipe/components/stream"
	"github.com/pipeproxy/pipe/internal/ctxcache"
	"github.com/wzshiming/logger"
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
		return ListenConfigNone
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
		store, ok := m.Load("stream.ListenConfig")
		if ok {
			o, ok := store.(map[string]stream.ListenConfig)[name]
			if ok {
				return o
			}
		}
	}

	if defaults != nil {
		return defaults
	}
	logger.FromContext(ctx).V(-1).Info("stream.ListenConfig is not defined", "name", name)
	return ListenConfigNone
}

var ListenConfigNone _ListenConfigNone

type _ListenConfigNone struct{}

func newListenConfigNone() stream.ListenConfig {
	return ListenConfigNone
}

func (_ListenConfigNone) IsVirtual() (_ bool) {
	logger.Log.V(-1).Info("this is none of stream.ListenConfig")

	return
}

func (_ListenConfigNone) ListenStream(_ context.Context) (_ net.Listener, error error) {
	logger.Log.V(-1).Info("this is none of stream.ListenConfig")

	error = fmt.Errorf("error stream.ListenConfig is none")

	return
}

type ListenConfig struct {
	Name string
	Def  stream.ListenConfig
	Ctx  context.Context
}

func (o *ListenConfig) IsVirtual() bool {
	return ListenConfigGet(o.Ctx, o.Name, o.Def).IsVirtual()
}

func (o *ListenConfig) ListenStream(context context.Context) (net.Listener, error) {
	return ListenConfigGet(o.Ctx, o.Name, o.Def).ListenStream(context)
}
