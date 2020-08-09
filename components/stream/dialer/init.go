package dialer

import (
	"fmt"

	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/components/stream"
	"github.com/wzshiming/pipe/components/tls"
)

const (
	name = "dialer"
)

var (
	ErrNotTLS = fmt.Errorf("not tls")
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
	Network DialerNetworkEnum
	Address string
	TLS     tls.TLS `json:",omitempty"`
}

func NewDialerWithConfig(conf *Config) stream.Dialer {
	if conf.TLS == nil {
		return NewDialer(string(conf.Network), conf.Address, nil)
	}
	return NewDialer(string(conf.Network), conf.Address, conf.TLS.TLS())
}
