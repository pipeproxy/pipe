// DO NOT EDIT! Code generated.
package tls

import (
	"context"

	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/components/tls"
	"github.com/wzshiming/pipe/internal/ctxcache"
	"github.com/wzshiming/pipe/internal/logger"
)

func init() {
	register.Register("ref", NewTLSRefWithConfig)
	register.Register("def", NewTLSDefWithConfig)
	register.Register("none", newTLSNone)
}

type Config struct {
	Name string
	Def  tls.TLS `json:",omitempty"`
}

func NewTLSRefWithConfig(ctx context.Context, conf *Config) tls.TLS {
	o := &TLS{
		Name: conf.Name,
		Def:  conf.Def,
		Ctx:  ctx,
	}
	return o
}

func NewTLSDefWithConfig(ctx context.Context, conf *Config) tls.TLS {
	return TLSPut(ctx, conf.Name, conf.Def)
}

func TLSPut(ctx context.Context, name string, def tls.TLS) tls.TLS {
	if def == nil {
		return TLSNone
	}

	m, ok := ctxcache.GetCacheWithContext(ctx)
	if !ok {
		return TLSNone
	}
	store, _ := m.LoadOrStore("tls.TLS", map[string]tls.TLS{})
	store.(map[string]tls.TLS)[name] = def
	return def
}

func TLSGet(ctx context.Context, name string, defaults tls.TLS) tls.TLS {
	m, ok := ctxcache.GetCacheWithContext(ctx)
	if ok {
		store, ok := m.Load("tls.TLS")
		if ok {
			o, ok := store.(map[string]tls.TLS)[name]
			if ok {
				return o
			}
		}
	}

	if defaults != nil {
		return defaults
	}
	logger.Warnf("tls.TLS %q is not defined", name)
	return TLSNone
}

var TLSNone _TLSNone

type _TLSNone struct{}

func newTLSNone() tls.TLS {
	return TLSNone
}

func (_TLSNone) TLS() (_ *tls.Config) {
	logger.Warn("this is none of tls.TLS")

	return
}

type TLS struct {
	Name string
	Def  tls.TLS
	Ctx  context.Context
}

func (o *TLS) TLS() *tls.Config {
	return TLSGet(o.Ctx, o.Name, o.Def).TLS()
}
