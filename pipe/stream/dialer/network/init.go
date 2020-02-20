package network

import (
	"github.com/wzshiming/pipe/configure/decode"
	"github.com/wzshiming/pipe/pipe/stream/dialer"
)

func init() {
	decode.Register(name, NewNetworkWithConfig)
}

const name = "network"

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
