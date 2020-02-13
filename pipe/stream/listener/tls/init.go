package tls

import (
	"github.com/wzshiming/pipe/configure/manager"
	"github.com/wzshiming/pipe/pipe/stream/listener"
	"github.com/wzshiming/pipe/pipe/tls"
)

const name = "tls"

func init() {
	manager.Register(name, NewTlsWithConfig)
}

type Config struct {
	ListenConfig listener.ListenConfig
	TLS          tls.TLS
}

func NewTlsWithConfig(conf *Config) listener.ListenConfig {
	return NewTls(conf.ListenConfig, conf.TLS.TLS())
}
