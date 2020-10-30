package listener

import (
	"context"

	"github.com/pipeproxy/pipe/components/packet"
	"github.com/pipeproxy/pipe/internal/listener"
	"github.com/pipeproxy/pipe/internal/logger"
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
