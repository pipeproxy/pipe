package tls

import (
	"crypto/tls"
)

type TLS interface {
	TLS() *tls.Config
}

type Config struct {
	tlsConfig *tls.Config
}

func NewConfig(tlsConfig *tls.Config) TLS {
	return &Config{tlsConfig: tlsConfig}
}

func (t *Config) TLS() *tls.Config {
	return t.tlsConfig
}
