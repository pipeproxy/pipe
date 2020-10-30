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
	ListenConfig stream.ListenConfig
	TLS          tls.TLS
}

func NewTlsWithConfig(conf *Config) stream.ListenConfig {
	return NewTls(conf.ListenConfig, conf.TLS)
}
