package network

import (
	"github.com/wzshiming/pipe/configure/manager"
	"github.com/wzshiming/pipe/pipe/stream/dialer"
)

func init() {
	manager.Register(name, NewNetworkWithConfig)
}

const name = "network"

type Config struct {
	Network string
	Address string
}

func NewNetworkWithConfig(conf *Config) dialer.Dialer {
	return NewNetwork(conf.Network, conf.Address)
}
