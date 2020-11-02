package listener

import (
	"context"
	"net"

	"github.com/gogf/greuse"
)

var (
	virtual      = newVirtualListenerManager()
	listenConfig = net.ListenConfig{
		Control: greuse.Control,
	}
	dialer = net.Dialer{}
)

func Listen(ctx context.Context, network, address string) (net.Listener, error) {
	return listenConfig.Listen(ctx, network, address)
}

func VirtualListen(ctx context.Context, network, address string) (net.Listener, error) {
	return virtual.Listen(ctx, network, address)
}

func ListenPacket(ctx context.Context, network, address string) (net.PacketConn, error) {
	return listenConfig.ListenPacket(ctx, network, address)
}

func DialContext(ctx context.Context, network, address string) (net.Conn, error) {
	return dialer.DialContext(ctx, network, address)
}

func VirtualDialContext(ctx context.Context, network, address string) (net.Conn, error) {
	return virtual.DialContext(ctx, network, address)
}
