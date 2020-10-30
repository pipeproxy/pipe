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

type ListenerNetworkEnum string

const (
	EnumUDP        ListenerNetworkEnum = "udp"
	EnumUDP4       ListenerNetworkEnum = "udp4"
	EnumUDP6       ListenerNetworkEnum = "udp6"
	EnumUnixPacket ListenerNetworkEnum = "unixpacket"
)

type Config struct {
	Network ListenerNetworkEnum
	Address string
}

func NewListenerWithConfig(ctx context.Context, conf *Config) packet.ListenConfig {
	return NewListener(string(conf.Network), conf.Address)
}
