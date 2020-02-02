package network

import (
	"context"
	"net"

	"github.com/wzshiming/pipe/internal/network"
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

func (n *Network) Listen(ctx context.Context) (net.Listener, error) {
	return network.Listen(ctx, n.network, n.address)
}
