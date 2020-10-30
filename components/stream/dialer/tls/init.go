package tls

import (
	"github.com/pipeproxy/pipe/components/common/register"
	"github.com/pipeproxy/pipe/components/stream"
	"github.com/pipeproxy/pipe/components/tls"
)

const (
	name = "tls"
)

func init() {
	register.Register(name, NewTlsWithConfig)
}

type Config struct {
	Dialer stream.Dialer
	TLS    tls.TLS
}

func NewTlsWithConfig(conf *Config) stream.Dialer {
	return NewTls(conf.Dialer, conf.TLS)
}
