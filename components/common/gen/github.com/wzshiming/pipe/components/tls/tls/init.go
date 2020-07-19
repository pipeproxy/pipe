// DO NOT EDIT! Code generated.
package reference

import (
	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/components/tls"
	"github.com/wzshiming/pipe/internal/logger"
)

func init() {
	register.Register("ref", NewTLSRefWithConfig)
	register.Register("def", NewTLSDefWithConfig)
	register.Register("none", NewTLSNone)
}

type Config struct {
	Name string
	Def  tls.TLS `json:",omitempty"`
}

func NewTLSRefWithConfig(conf *Config) (tls.TLS, error) {
	o := &TLS{
		Name: conf.Name,
		Def:  conf.Def,
	}
	return o, nil
}

func NewTLSDefWithConfig(conf *Config) (tls.TLS, error) {
	TLSStore[conf.Name] = conf.Def
	return conf.Def, nil
}

var TLSStore = map[string]tls.TLS{}

func TLSFind(name string, defaults tls.TLS) tls.TLS {
	o, ok := TLSStore[name]
	if ok {
		return o
	}
	if defaults != nil {
		return defaults
	}
	return TLSNone{}
}

type TLSNone struct{}

func NewTLSNone() tls.TLS {
	return TLSNone{}
}

func (TLSNone) TLS() (_ *tls.Config) {
	logger.Warn("this is none of tls.TLS")
	return
}

type TLS struct {
	Name string
	Def  tls.TLS
}

func (o *TLS) TLS() *tls.Config {
	return TLSFind(o.Name, o.Def).TLS()
}
