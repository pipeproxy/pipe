package packet

import (
	"github.com/pipeproxy/pipe/components/common/register"
	"github.com/pipeproxy/pipe/components/packet"
	"github.com/pipeproxy/pipe/components/service"
)

const (
	name = "packet"
)

func init() {
	register.Register(name, NewServerWithConfig)
}

type Config struct {
	Listener packet.ListenConfig
	Handler  packet.Handler
}

func NewServerWithConfig(conf *Config) (service.Service, error) {
	return NewServer(conf.Listener, conf.Handler)
}
