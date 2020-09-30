package listener

import (
	"context"

	"github.com/wzshiming/pipe/components/stream"
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

func (n *Listener) ListenStream(ctx context.Context) (stream.StreamListener, error) {
	return listener.Listen(ctx, n.network, n.address)
}
