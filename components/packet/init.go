package packet

import (
	"context"
	"net"

	"github.com/wzshiming/pipe/components/common/types"
)

func init() {
	var handler Handler
	types.Register(&handler)
}

type Packet = net.PacketConn

type Handler interface {
	ServePacket(ctx context.Context, pkt Packet)
}
