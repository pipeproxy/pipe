package tls

import (
	"github.com/wzshiming/pipe/configure"
	"github.com/wzshiming/pipe/stream"
	"github.com/wzshiming/pipe/tls"
)

const nameUp = "tls_up"
const nameDown = "tls_down"

func init() {
	configure.Register(nameUp, NewTlsUpWithConfig)
	configure.Register(nameDown, NewTlsDownWithConfig)
}

type Config struct {
	Handler stream.Handler
	TLS     tls.TLS
}

func NewTlsDownWithConfig(conf *Config) stream.Handler {
	return NewTlsDown(conf.Handler, conf.TLS.TLS())
}

func NewTlsUpWithConfig(conf *Config) stream.Handler {
	return NewTlsUp(conf.Handler, conf.TLS.TLS())
}
