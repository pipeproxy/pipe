package tls

import (
	"crypto/tls"

	"github.com/wzshiming/pipe/configure/alias"
)

func init() {
	var tls TLS
	alias.Register("TLS", &tls)
}

type Config = tls.Config

type TLS interface {
	TLS() *Config
}

type wrapTLS struct {
	tlsConfig *Config
}

func WrapTLS(tlsConfig *Config) TLS {
	return &wrapTLS{tlsConfig: tlsConfig}
}

func (t *wrapTLS) TLS() *Config {
	return t.tlsConfig
}
