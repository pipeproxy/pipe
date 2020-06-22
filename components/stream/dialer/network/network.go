package network

import (
	"context"
	"net"

	"github.com/wzshiming/pipe/components/stream"
	"github.com/wzshiming/pipe/internal/logger"
)

type Network struct {
	network string
	address string
	dialer  net.Dialer
}

func NewNetwork(network string, address string) *Network {
	return &Network{
		network: network,
		address: address,
	}
}

func (n *Network) DialStream(ctx context.Context) (stream.Stream, error) {
	stm, err := n.dialer.DialContext(ctx, n.network, n.address)
	if err != nil {
		logger.Errorf("Forward error: %s", err)
	}
	return stm, err
}
