package packet

import (
	"context"
	"net"
)

type Packet = net.PacketConn

type Handler interface {
	ServePacket(ctx context.Context, pkt Packet)
}
