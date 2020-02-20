package packet

import (
	"context"
	"net"

	"github.com/wzshiming/pipe/configure/alias"
)

func init() {
	var handler Handler
	alias.Register("packet.Handler", &handler)
}

type Packet = net.PacketConn

type Handler interface {
	ServePacket(ctx context.Context, pkt Packet)
}
