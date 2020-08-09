package listener

import (
	"context"
	"crypto/tls"

	"github.com/wzshiming/pipe/components/stream"
	"github.com/wzshiming/pipe/internal/listener"
)

type Listener struct {
	network   string
	address   string
	tlsConfig *tls.Config
}

func NewListener(network string, address string, tlsConfig *tls.Config) *Listener {
	return &Listener{
		network:   network,
		address:   address,
		tlsConfig: tlsConfig,
	}
}

func (n *Listener) ListenStream(ctx context.Context) (stream.StreamListener, error) {
	listen, err := listener.Listen(ctx, n.network, n.address)
	if err != nil {
		return nil, err
	}
	if n.tlsConfig != nil {
		listen = tls.NewListener(listen, n.tlsConfig)
	}
	return listen, nil
}
