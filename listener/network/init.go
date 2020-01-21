package network

import (
	"context"

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

func NewNetworkWithConfig(ctx context.Context, conf *Config) listener.ListenConfig {
	return NewNetwork(conf.Network, conf.Address)
}
