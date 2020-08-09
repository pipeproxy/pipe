package tls

import (
	"crypto/tls"
)

type wrapTLS struct {
	tlsConfig *Config
}

func WrapTLS(tlsConfig *Config) TLS {
	return &wrapTLS{tlsConfig: tlsConfig}
}

func (t *wrapTLS) TLS() *Config {
	return t.tlsConfig
}

var (
	Client      = tls.Client
	Server      = tls.Server
	NewListener = tls.NewListener
)
