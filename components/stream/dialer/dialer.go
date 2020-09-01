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

func (d *Dialer) DialStream(ctx context.Context) (stream.Stream, error) {
	stm, err := d.dialer.DialContext(ctx, d.network, d.address)
	if err != nil {
		return nil, err
	}
	if d.tlsConfig != nil {
		stm = tls.Client(stm, d.tlsConfig)
	}
	return stm, nil
}

func (d *Dialer) IsTCP() bool {
	return d.tlsConfig != nil
}
