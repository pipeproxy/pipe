package network

import (
	"context"
	"log"
	"net"

	"github.com/wzshiming/pipe/configure"
	"github.com/wzshiming/pipe/listener"
)

func init() {
	configure.Register(name, NewNetworkWithConfig)
}

const name = "network"

type Config struct {
	Network string
	Address string
}

var ListenConfig net.ListenConfig

func NewNetworkWithConfig(ctx context.Context, conf *Config) (listener.Listener, error) {
	log.Printf("[INFO] Listen to %s://%s", conf.Network, conf.Address)
	return ListenConfig.Listen(ctx, conf.Network, conf.Address)
}
