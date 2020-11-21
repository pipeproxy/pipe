package listener

import (
	"context"

	"github.com/pipeproxy/pipe/components/stream"
	"github.com/pipeproxy/pipe/internal/listener"
	"github.com/wzshiming/logger"
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
	log := logger.FromContext(ctx)
	if l.virtual {
		log.Info("Virtual listen stream",
			"localAddress", l.address,
			"virtual", true,
		)
		return listener.VirtualListen(ctx, l.network, l.address)
	}
	log.Info("Listen stream",
		"localAddress", l.address,
	)
	return listener.Listen(ctx, l.network, l.address)
}
