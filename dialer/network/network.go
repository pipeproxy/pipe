package network

import (
	"context"
	"log"
	"net"

	"github.com/wzshiming/pipe/stream"
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

func (n *Network) Dial(ctx context.Context) (stream.Stream, error) {
	stm, err := n.dialer.DialContext(ctx, n.network, n.address)
	if err != nil {
		log.Printf("[ERROR] Forward error: %s", err.Error())
	}
	return stm, err
}
