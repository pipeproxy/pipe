package tls

import (
	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/components/stream/dialer"
	"github.com/wzshiming/pipe/components/tls"
)

const name = "tls"

func init() {
	register.Register(name, NewTlsWithConfig)
}

type Config struct {
	Dialer dialer.Dialer
	TLS    tls.TLS
}

func NewTlsWithConfig(conf *Config) dialer.Dialer {
	return NewTls(conf.Dialer, conf.TLS.TLS())
}
