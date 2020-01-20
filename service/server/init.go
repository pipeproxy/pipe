package server

import (
	"github.com/wzshiming/pipe/configure"
	"github.com/wzshiming/pipe/listener"
	"github.com/wzshiming/pipe/service"
	"github.com/wzshiming/pipe/stream"
)

const name = "server"

func init() {
	configure.Register(name, NewServerWithConfig)
}

type Config struct {
	Listener listener.Listener
	Handler  stream.Handler
}

func NewServerWithConfig(conf *Config) service.Service {
	return NewServer(conf.Listener, conf.Handler)
}
