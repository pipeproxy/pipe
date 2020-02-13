package tls

import (
	"crypto/tls"
)

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
