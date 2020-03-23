package network

import (
	"context"

	"github.com/wzshiming/pipe/components/stream/listener"
	"github.com/wzshiming/pipe/internal/stream"
)

type Network struct {
	network string
	address string
}

func NewNetwork(network string, address string) *Network {
	return &Network{
		network: network,
		address: address,
	}
}

func (n *Network) ListenStream(ctx context.Context) (listener.StreamListener, error) {
	return stream.Listen(ctx, n.network, n.address)
}
