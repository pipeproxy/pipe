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
	if log.Enabled() {
		if l.virtual {
			log.Info("Virtual listen stream",
				"localAddress", l.address,
				"virtual", true,
			)
		} else {
			log.Info("Listen stream",
				"localAddress", l.address,
			)
		}
	}
	if l.virtual {
		return listener.VirtualListen(ctx, l.network, l.address)
	}
	return listener.Listen(ctx, l.network, l.address)
}

func (l *Listener) IsVirtual() bool {
	return l.virtual
}
