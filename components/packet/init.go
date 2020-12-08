package packet

import (
	"context"
	"net"

	"github.com/pipeproxy/pipe/components/common/types"
)

func init() {
	var handler Handler
	types.Register(&handler)
	var listenConfig ListenConfig
	types.Register(&listenConfig)
	var packet Packet
	types.Register(&packet)
}

type Packet = net.PacketConn

type Handler interface {
	ServePacket(ctx context.Context, pkt Packet)
}

type ListenConfig interface {
	ListenPacket(ctx context.Context) (Packet, error)
}

type NetworkEnum string

const (
	EnumNetworkUDP        NetworkEnum = "udp"
	EnumNetworkUDP4       NetworkEnum = "udp4"
	EnumNetworkUDP6       NetworkEnum = "udp6"
	EnumNetworkUnixPacket NetworkEnum = "unixpacket"
)
