package listener

import (
	"context"

	"github.com/pipeproxy/pipe/components/common/register"
	"github.com/pipeproxy/pipe/components/stream"
)

const (
	name = "listener"
)

func init() {
	register.Register(name, NewListenerWithConfig)
}

type ListenerNetworkEnum string

const (
	EnumTCP  ListenerNetworkEnum = "tcp"
	EnumTCP4 ListenerNetworkEnum = "tcp4"
	EnumTCP6 ListenerNetworkEnum = "tcp6"
	EnumUnix ListenerNetworkEnum = "unix"
)

type Config struct {
	Network ListenerNetworkEnum
	Address string
	Virtual bool `json:",omitempty"`
}

func NewListenerWithConfig(ctx context.Context, conf *Config) stream.ListenConfig {
	return NewListener(string(conf.Network), conf.Address, conf.Virtual)
}
