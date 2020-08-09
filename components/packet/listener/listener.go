package listener

import (
	"context"

	"github.com/wzshiming/pipe/components/packet"
	"github.com/wzshiming/pipe/internal/listener"
)

type Listener struct {
	network string
	address string
}

func NewListener(network string, address string) *Listener {
	return &Listener{
		network: network,
		address: address,
	}
}

func (n *Listener) ListenPacket(ctx context.Context) (packet.Packet, error) {
	return listener.ListenPacket(ctx, n.network, n.address)
}
