// DO NOT EDIT! Code generated.
package dialer

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
	register.Register("ref", NewDialerRefWithConfig)
	register.Register("def", NewDialerDefWithConfig)
	register.Register("none", newDialerNone)
}

type Config struct {
	Name string
	Def  stream.Dialer `json:",omitempty"`
}

func NewDialerRefWithConfig(ctx context.Context, conf *Config) stream.Dialer {
	o := &Dialer{
		Name: conf.Name,
		Def:  conf.Def,
		Ctx:  ctx,
	}
	return o
}

func NewDialerDefWithConfig(ctx context.Context, conf *Config) stream.Dialer {
	return DialerPut(ctx, conf.Name, conf.Def)
}

func DialerPut(ctx context.Context, name string, def stream.Dialer) stream.Dialer {
	if def == nil {
		return DialerNone
	}

	m, ok := ctxcache.GetCacheWithContext(ctx)
	if !ok {
		return DialerNone
	}
	store, _ := m.LoadOrStore("dialer.Dialer", map[string]stream.Dialer{})
	store.(map[string]stream.Dialer)[name] = def
	return def
}

func DialerGet(ctx context.Context, name string, defaults stream.Dialer) stream.Dialer {
	m, ok := ctxcache.GetCacheWithContext(ctx)
	if ok {
		store, ok := m.Load("dialer.Dialer")
		if ok {
			o, ok := store.(map[string]stream.Dialer)[name]
			if ok {
				return o
			}
		}
	}

	if defaults != nil {
		return defaults
	}
	logger.Warnf("dialer.Dialer %q is not defined", name)
	return DialerNone
}

var DialerNone _DialerNone

type _DialerNone struct{}

func newDialerNone() stream.Dialer {
	return DialerNone
}

func (_DialerNone) DialStream(_ context.Context) (_ net.Conn, error error) {
	logger.Warn("this is none of dialer.Dialer")

	error = fmt.Errorf("error dialer.Dialer is none")

	return
}

type Dialer struct {
	Name string
	Def  stream.Dialer
	Ctx  context.Context
}

func (o *Dialer) DialStream(context context.Context) (net.Conn, error) {
	return DialerGet(o.Ctx, o.Name, o.Def).DialStream(context)
}
