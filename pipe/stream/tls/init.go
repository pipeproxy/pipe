package tls

import (
	"github.com/wzshiming/pipe/configure/decode"
	"github.com/wzshiming/pipe/pipe/stream"
	"github.com/wzshiming/pipe/pipe/tls"
)

const nameUp = "tls_up"
const nameDown = "tls_down"

func init() {
	decode.Register(nameUp, NewTlsUpWithConfig)
	decode.Register(nameDown, NewTlsDownWithConfig)
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
