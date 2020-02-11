package poller

import (
	"fmt"
	"net/http"

	"github.com/wzshiming/pipe/configure/manager"
)

var (
	ErrNotHandler = fmt.Errorf("error not handler")
	ErrNotRoller  = fmt.Errorf("error not poller")
)

const name = "poller"

func init() {
	manager.Register(name, NewPollerWithConfig)
}

type Config struct {
	Poller   string
	Handlers []http.Handler
}

func NewPollerWithConfig(conf *Config) (http.Handler, error) {
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
