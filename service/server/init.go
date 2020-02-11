package server

import (
	"github.com/wzshiming/pipe/configure/manager"
	"github.com/wzshiming/pipe/listener"
	"github.com/wzshiming/pipe/service"
	"github.com/wzshiming/pipe/stream"
)

const name = "server"

func init() {
	manager.Register(name, NewServerWithConfig)
}

type Config struct {
	Listener listener.ListenConfig
	Handler  stream.Handler
}

func NewServerWithConfig(conf *Config) (service.Service, error) {
	return NewServer(conf.Listener, conf.Handler)
}
