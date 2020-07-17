package poller

import (
	"fmt"

	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/components/stream/dialer"
)

var (
	ErrNotDialer = fmt.Errorf("error not dialer")
	ErrNotRoller = fmt.Errorf("error not poller")
)

const (
	name = "poller"
)

func init() {
	register.Register(name, NewPollerWithConfig)
}

type PollerEnum string

const (
	EnumRandom     PollerEnum = "random"
	EnumRoundRobin PollerEnum = "round_robin"
)

type Config struct {
	Poller  PollerEnum
	Dialers []dialer.Dialer
}

func NewPollerWithConfig(conf *Config) (dialer.Dialer, error) {
	switch len(conf.Dialers) {
	case 0:
		return nil, ErrNotDialer
	case 1:
		return conf.Dialers[0], nil
	}

	switch conf.Poller {
	case EnumRandom:
		return NewRandom(conf.Dialers), nil
	case EnumRoundRobin:
		return NewRoundRobin(conf.Dialers), nil
	default:
		return nil, ErrNotRoller
	}
}
