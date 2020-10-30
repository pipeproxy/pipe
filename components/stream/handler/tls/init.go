package tls

import (
	"github.com/pipeproxy/pipe/components/common/register"
	"github.com/pipeproxy/pipe/components/stream"
	"github.com/pipeproxy/pipe/components/tls"
)

const (
	nameUp   = "tls_up"
	nameDown = "tls_down"
)

func init() {
	register.Register(nameUp, NewTlsUpWithConfig)
	register.Register(nameDown, NewTlsDownWithConfig)
}

type Config struct {
	Handler stream.Handler
	TLS     tls.TLS
}

func NewTlsDownWithConfig(conf *Config) stream.Handler {
	return NewTlsDown(conf.Handler, conf.TLS)
}

func NewTlsUpWithConfig(conf *Config) stream.Handler {
	return NewTlsUp(conf.Handler, conf.TLS)
}
