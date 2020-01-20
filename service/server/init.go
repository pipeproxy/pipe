package server

import (
	"context"

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
	Handlers []stream.Handler
}

func NewServerWithConfig(ctx context.Context, config []byte) (service.Service, error) {
	var conf Config
	err := configure.Decode(ctx, config, &conf)
	if err != nil {
		return nil, err
	}
	return NewServer(conf.Listener, conf.Handlers)
}
