package poller

import (
	"fmt"

	"github.com/wzshiming/pipe/configure/decode"
	"github.com/wzshiming/pipe/pipe/stream/dialer"
)

var (
	ErrNotDialer = fmt.Errorf("error not dialer")
	ErrNotRoller = fmt.Errorf("error not poller")
)

const name = "poller"

func init() {
	decode.Register(name, NewPollerWithConfig)
}

type Config struct {
	Poller  string
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
	case "random":
		return NewRandom(conf.Dialers), nil
	case "round_robin":
		return NewRoundRobin(conf.Dialers), nil
	default:
		return nil, ErrNotRoller
	}
}
