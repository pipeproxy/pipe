package listener

import (
	"context"

	"github.com/pipeproxy/pipe/components/common/register"
	"github.com/pipeproxy/pipe/components/packet"
)

const (
	name = "listener"
)

func init() {
	register.Register(name, NewListenerWithConfig)
}

type Config struct {
	Network packet.NetworkEnum
	Address string
}

func NewListenerWithConfig(ctx context.Context, conf *Config) packet.ListenConfig {
	return NewListener(string(conf.Network), conf.Address)
}
