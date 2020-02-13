package tls

import (
	"github.com/wzshiming/pipe/configure/manager"
	"github.com/wzshiming/pipe/pipe/stream/dialer"
	"github.com/wzshiming/pipe/pipe/tls"
)

const name = "tls"

func init() {
	manager.Register(name, NewTlsWithConfig)
}

type Config struct {
	Dialer dialer.Dialer
	TLS    tls.TLS
}

func NewTlsWithConfig(conf *Config) dialer.Dialer {
	return NewTls(conf.Dialer, conf.TLS.TLS())
}
