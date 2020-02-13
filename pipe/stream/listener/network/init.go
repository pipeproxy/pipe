package network

import (
	"context"

	"github.com/wzshiming/pipe/configure/manager"
	"github.com/wzshiming/pipe/pipe/stream/listener"
)

func init() {
	manager.Register(name, NewNetworkWithConfig)
}

const name = "network"

type Config struct {
	Network string
	Address string
}

func NewNetworkWithConfig(ctx context.Context, conf *Config) listener.ListenConfig {
	return NewNetwork(conf.Network, conf.Address)
}
