package tls

import (
	"crypto/tls"

	"github.com/pipeproxy/pipe/components/common/types"
)

func init() {
	var tls TLS
	types.Register(&tls)
}

type Config = tls.Config

type TLS interface {
	TLS() *Config
}
