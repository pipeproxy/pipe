package tls

import (
	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/components/stream"
	"github.com/wzshiming/pipe/components/tls"
)

const (
	name = "tls"
)

func init() {
	register.Register(name, NewTlsWithConfig)
}

type Config struct {
	Dialer stream.Dialer
	TLS    tls.TLS `json:",omitempty"`
}

func NewTlsWithConfig(conf *Config) stream.Dialer {
	return NewTls(conf.Dialer, conf.TLS.TLS())
}
