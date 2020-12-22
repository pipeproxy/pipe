package listener

import (
	"context"

	"github.com/pipeproxy/pipe/components/packet"
	"github.com/pipeproxy/pipe/internal/listener"
	"github.com/wzshiming/logger"
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
	log := logger.FromContext(ctx)
	if log.Enabled() {
		log.Info("Listen packet",
			"localAddress", n.address,
		)
	}
	return listener.ListenPacket(ctx, n.network, n.address)
}
