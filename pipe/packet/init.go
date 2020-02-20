package packet

import (
	"context"
	"net"

	"github.com/wzshiming/pipe/configure/alias"
	"github.com/wzshiming/pipe/pipe/common/load"
)

func init() {
	var handler Handler
	alias.Register("packet.Handler", &handler)
	load.Register(&handler)
}

type Packet = net.PacketConn

type Handler interface {
	ServePacket(ctx context.Context, pkt Packet)
}
