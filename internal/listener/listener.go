package listener

import (
	"context"
	"net"

	"github.com/gogf/greuse"
)

var listenConfig = net.ListenConfig{
	Control: greuse.Control,
}

func Listen(ctx context.Context, network, address string) (net.Listener, error) {
	return listenConfig.Listen(ctx, network, address)
}

func ListenPacket(ctx context.Context, network, address string) (net.PacketConn, error) {
	return listenConfig.ListenPacket(ctx, network, address)
}
