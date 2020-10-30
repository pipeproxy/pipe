package listener

import (
	"context"

	"github.com/pipeproxy/pipe/components/stream"
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

func (n *Listener) ListenStream(ctx context.Context) (stream.StreamListener, error) {
	logger.Infof("Listen %s://%s", n.network, n.address)
	return listener.Listen(ctx, n.network, n.address)
}
