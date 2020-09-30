package dialer

import (
	"context"
	"net"

	"github.com/wzshiming/pipe/components/stream"
)

type Dialer struct {
	network string
	address string
	dialer  net.Dialer
}

func NewDialer(network string, address string) *Dialer {
	return &Dialer{
		network: network,
		address: address,
	}
}

func (d *Dialer) DialStream(ctx context.Context) (stream.Stream, error) {
	return d.dialer.DialContext(ctx, d.network, d.address)
}
