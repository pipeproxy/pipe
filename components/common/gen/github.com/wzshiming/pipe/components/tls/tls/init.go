// DO NOT EDIT! Code generated.
package reference

import (
	"sync"

	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/components/tls"
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

func NewTLSRefWithConfig(conf *Config) tls.TLS {
	o := &TLS{
		Name: conf.Name,
		Def:  conf.Def,
	}
	return o
}

func NewTLSDefWithConfig(conf *Config) tls.TLS {
	return TLSPut(conf.Name, conf.Def)
}

var (
	mut       sync.RWMutex
	_TLSStore = map[string]tls.TLS{}
)

func TLSPut(name string, def tls.TLS) tls.TLS {
	if def == nil {
		def = TLSNone
	}
	mut.Lock()
	_TLSStore[name] = def
	mut.Unlock()
	return def
}

func TLSGet(name string, defaults tls.TLS) tls.TLS {
	mut.RLock()
	o, ok := _TLSStore[name]
	mut.RUnlock()
	if ok {
		return o
	}
	if defaults != nil {
		return defaults
	}
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
}

func (o *TLS) TLS() *tls.Config {
	return TLSGet(o.Name, o.Def).TLS()
}
