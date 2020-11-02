package dialer

import (
	"github.com/pipeproxy/pipe/components/common/register"
	"github.com/pipeproxy/pipe/components/stream"
)

const (
	name = "dialer"
)

func init() {
	register.Register(name, NewDialerWithConfig)
}

type DialerNetworkEnum string

const (
	EnumTCP  DialerNetworkEnum = "tcp"
	EnumTCP4 DialerNetworkEnum = "tcp4"
	EnumTCP6 DialerNetworkEnum = "tcp6"
	EnumUnix DialerNetworkEnum = "unix"
)

type Config struct {
	Network  DialerNetworkEnum
	Address  string
	Original bool `json:",omitempty"`
	Virtual  bool `json:",omitempty"`
}

func NewDialerWithConfig(conf *Config) stream.Dialer {
	return NewDialer(string(conf.Network), conf.Address, conf.Virtual, conf.Original)
}
