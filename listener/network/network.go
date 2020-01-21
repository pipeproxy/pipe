package network

import (
	"context"
	"log"
	"net"
)

type Network struct {
	network      string
	address      string
	listenConfig net.ListenConfig
}

func NewNetwork(network string, address string) *Network {
	return &Network{
		network: network,
		address: address,
	}
}

func (n *Network) Listen(ctx context.Context) (net.Listener, error) {
	log.Printf("[INFO] Listen to %s://%s", n.network, n.address)
	return n.listenConfig.Listen(ctx, n.network, n.address)
}
