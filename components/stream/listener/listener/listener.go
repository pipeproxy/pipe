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
	virtual bool
}

func NewListener(network string, address string, virtual bool) *Listener {
	return &Listener{
		network: network,
		address: address,
		virtual: virtual,
	}
}

func (l *Listener) ListenStream(ctx context.Context) (stream.StreamListener, error) {
	if l.virtual {
		logger.Infof("Virtual Listen %s://%s", l.network, l.address)
		return listener.VirtualListen(ctx, l.network, l.address)
	}
	logger.Infof("Listen %s://%s", l.network, l.address)
	return listener.Listen(ctx, l.network, l.address)
}
