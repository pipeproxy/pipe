package listener

import (
	"context"

	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/components/stream"
	"github.com/wzshiming/pipe/components/tls"
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
	TLS     tls.TLS `json:",omitempty"`
}

func NewListenerWithConfig(ctx context.Context, conf *Config) stream.ListenConfig {
	if conf.TLS == nil {
		return NewListener(string(conf.Network), conf.Address, nil)
	}
	return NewListener(string(conf.Network), conf.Address, conf.TLS.TLS())
}
