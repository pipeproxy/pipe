package tls

import (
	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/components/stream/listener"
	"github.com/wzshiming/pipe/components/tls"
)

const (
	name = "tls"
)

func init() {
	register.Register(name, NewTlsWithConfig)
}

type Config struct {
	ListenConfig listener.ListenConfig
	TLS          tls.TLS
}

func NewTlsWithConfig(conf *Config) listener.ListenConfig {
	return NewTls(conf.ListenConfig, conf.TLS.TLS())
}
