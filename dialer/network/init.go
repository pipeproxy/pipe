package network

import (
	"github.com/wzshiming/pipe/configure"
	"github.com/wzshiming/pipe/dialer"
)

func init() {
	configure.Register(name, NewNetworkWithConfig)
}

const name = "network"

type Config struct {
	Network string
	Address string
}

func NewNetworkWithConfig(conf *Config) dialer.Dialer {
	return NewNetwork(conf.Network, conf.Address)
}
