package listener

import (
	"context"

	"github.com/wzshiming/pipe/components/packet"
	"github.com/wzshiming/pipe/internal/listener"
	"github.com/wzshiming/pipe/internal/logger"
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
	logger.Infof("Listen %s://%s", n.network, n.address)
	return listener.ListenPacket(ctx, n.network, n.address)
}
