// DO NOT EDIT! Code generated.
package listener

import (
	"context"
	"fmt"
	"net"

	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/internal/ctxcache"
	"github.com/wzshiming/pipe/internal/logger"
)

func init() {
	register.Register("ref", NewListenerRefWithConfig)
	register.Register("def", NewListenerDefWithConfig)
	register.Register("none", newListenerNone)
}

type Config struct {
	Name string
	Def  net.Listener `json:",omitempty"`
}

func NewListenerRefWithConfig(ctx context.Context, conf *Config) net.Listener {
	o := &Listener{
		Name: conf.Name,
		Def:  conf.Def,
		Ctx:  ctx,
	}
	return o
}

func NewListenerDefWithConfig(ctx context.Context, conf *Config) net.Listener {
	return ListenerPut(ctx, conf.Name, conf.Def)
}

func ListenerPut(ctx context.Context, name string, def net.Listener) net.Listener {
	if def == nil {
		def = ListenerNone
	}

	m, ok := ctxcache.GetCacheWithContext(ctx)
	if !ok {
		return ListenerNone
	}
	store, _ := m.LoadOrStore("net.Listener", map[string]net.Listener{})
	store.(map[string]net.Listener)[name] = def
	return def
}

func ListenerGet(ctx context.Context, name string, defaults net.Listener) net.Listener {
	m, ok := ctxcache.GetCacheWithContext(ctx)
	if ok {
		store, _ := m.LoadOrStore("net.Listener", map[string]net.Listener{})
		o, ok := store.(map[string]net.Listener)[name]
		if ok {
			return o
		}
	}

	if defaults != nil {
		return defaults
	}
	return ListenerNone
}

var ListenerNone _ListenerNone

type _ListenerNone struct{}

func newListenerNone() net.Listener {
	return ListenerNone
}

func (_ListenerNone) Accept() (_ net.Conn, error error) {
	logger.Warn("this is none of net.Listener")

	error = fmt.Errorf("error none")

	return
}

func (_ListenerNone) Addr() (_ net.Addr) {
	logger.Warn("this is none of net.Listener")

	return
}

func (_ListenerNone) Close() (error error) {
	logger.Warn("this is none of net.Listener")

	error = fmt.Errorf("error none")

	return
}

type Listener struct {
	Name string
	Def  net.Listener
	Ctx  context.Context
}

func (o *Listener) Accept() (net.Conn, error) {
	return ListenerGet(o.Ctx, o.Name, o.Def).Accept()
}

func (o *Listener) Addr() net.Addr {
	return ListenerGet(o.Ctx, o.Name, o.Def).Addr()
}

func (o *Listener) Close() error {
	return ListenerGet(o.Ctx, o.Name, o.Def).Close()
}
