package stream

import (
	"github.com/wzshiming/pipe/configure/decode"
	"github.com/wzshiming/pipe/pipe/service"
	"github.com/wzshiming/pipe/pipe/stream"
	"github.com/wzshiming/pipe/pipe/stream/listener"
)

const name = "stream"

func init() {
	decode.Register(name, NewServerWithConfig)
}

type Config struct {
	Listener listener.ListenConfig
	Handler  stream.Handler
}

func NewServerWithConfig(conf *Config) (service.Service, error) {
	return NewServer(conf.Listener, conf.Handler)
}
