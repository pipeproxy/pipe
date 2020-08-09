package stream

import (
	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/components/service"
	"github.com/wzshiming/pipe/components/stream"
)

const (
	name = "stream"
)

func init() {
	register.Register(name, NewServerWithConfig)
}

type Config struct {
	Listener          stream.ListenConfig
	Handler           stream.Handler
	DisconnectOnClose bool `json:",omitempty"`
}

func NewServerWithConfig(conf *Config) (service.Service, error) {
	return NewServer(conf.Listener, conf.Handler, conf.DisconnectOnClose)
}
