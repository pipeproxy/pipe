package network

import (
	"context"

	"github.com/wzshiming/pipe/configure/decode"
	"github.com/wzshiming/pipe/pipe/stream/listener"
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

func NewNetworkWithConfig(ctx context.Context, conf *Config) listener.ListenConfig {
	return NewNetwork(string(conf.Network), conf.Address)
}
