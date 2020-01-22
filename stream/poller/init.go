package poller

import (
	"fmt"

	"github.com/wzshiming/pipe/configure"
	"github.com/wzshiming/pipe/stream"
)

var (
	ErrNotHandler = fmt.Errorf("error not handler")
	ErrNotRoller  = fmt.Errorf("error not poller")
)

const name = "poller"

func init() {
	configure.Register(name, NewPollerWithConfig)
}

type Config struct {
	Poller   string
	Handlers []stream.Handler
}

func NewPollerWithConfig(conf *Config) (stream.Handler, error) {
	if len(conf.Handlers) == 0 {
		return nil, ErrNotHandler
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
