package tls

import (
	"crypto/tls"

	"github.com/wzshiming/pipe/configure/alias"
	"github.com/wzshiming/pipe/pipe/common/load"
)

func init() {
	var tls TLS
	alias.Register("TLS", &tls)
	load.Register(&tls)
}

type Config = tls.Config

type TLS interface {
	TLS() *Config
}
