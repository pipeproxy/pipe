package tls

import (
	"github.com/wzshiming/pipe/configure/decode"
	"github.com/wzshiming/pipe/pipe/stream/listener"
	"github.com/wzshiming/pipe/pipe/tls"
)

const name = "tls"

func init() {
	decode.Register(name, NewTlsWithConfig)
}

type Config struct {
	ListenConfig listener.ListenConfig
	TLS          tls.TLS
}

func NewTlsWithConfig(conf *Config) listener.ListenConfig {
	return NewTls(conf.ListenConfig, conf.TLS.TLS())
}
