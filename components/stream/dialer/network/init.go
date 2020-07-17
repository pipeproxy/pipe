package network

import (
	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/components/stream/dialer"
)

const (
	name = "network"
)

func init() {
	register.Register(name, NewNetworkWithConfig)
}

type NetworkEnum string

const (
	EnumTCP  NetworkEnum = "tcp"
	EnumTCP4 NetworkEnum = "tcp4"
	EnumTCP6 NetworkEnum = "tcp6"
	EnumUnix NetworkEnum = "unix"
)

type Config struct {
	Network NetworkEnum
	Address string
}

func NewNetworkWithConfig(conf *Config) dialer.Dialer {
	return NewNetwork(string(conf.Network), conf.Address)
}
