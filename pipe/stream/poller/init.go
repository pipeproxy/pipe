package poller

import (
	"fmt"

	"github.com/wzshiming/pipe/configure/decode"
	"github.com/wzshiming/pipe/pipe/stream"
)

var (
	ErrNotHandler = fmt.Errorf("error not handler")
	ErrNotRoller  = fmt.Errorf("error not poller")
)

const name = "poller"

func init() {
	decode.Register(name, NewPollerWithConfig)
}

type Config struct {
	Poller   string
	Handlers []stream.Handler
}

func NewPollerWithConfig(conf *Config) (stream.Handler, error) {
	switch len(conf.Handlers) {
	case 0:
		return nil, ErrNotHandler
	case 1:
		return conf.Handlers[0], nil
	}

	switch conf.Poller {
	case "random":
		return NewRandom(conf.Handlers), nil
	case "round_robin":
		return NewRoundRobin(conf.Handlers), nil
	default:
		return nil, ErrNotRoller
	}
}
