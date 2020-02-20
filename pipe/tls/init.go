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
