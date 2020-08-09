package dialer

import (
	"context"
	"net"

	"github.com/wzshiming/pipe/components/stream"
	"github.com/wzshiming/pipe/components/tls"
)

type Dialer struct {
	network   string
	address   string
	dialer    net.Dialer
	tlsConfig *tls.Config
}

func NewDialer(network string, address string, tlsConfig *tls.Config) *Dialer {
	return &Dialer{
		network:   network,
		address:   address,
		tlsConfig: tlsConfig,
	}
}

func (n *Dialer) DialStream(ctx context.Context) (stream.Stream, error) {
	stm, err := n.dialer.DialContext(ctx, n.network, n.address)
	if err != nil {
		return nil, err
	}
	if n.tlsConfig != nil {
		stm = tls.Client(stm, n.tlsConfig)
	}
	return stm, nil
}
