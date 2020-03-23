package poller

import (
	"fmt"
	"net/http"

	"github.com/wzshiming/pipe/components/common/register"
)

var (
	ErrNotHandler = fmt.Errorf("error not handler")
	ErrNotRoller  = fmt.Errorf("error not poller")
)

const name = "poller"

func init() {
	register.Register(name, NewPollerWithConfig)
}

type PollerEnum string

const (
	EnumRandom     PollerEnum = "random"
	EnumRoundRobin PollerEnum = "round_robin"
)

type Config struct {
	Poller   PollerEnum
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
	case EnumRandom:
		return NewRandom(conf.Handlers), nil
	case EnumRoundRobin:
		return NewRoundRobin(conf.Handlers), nil
	default:
		return nil, ErrNotRoller
	}
}
