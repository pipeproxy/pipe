package lb

import (
	"fmt"
	"net/http"

	"github.com/pipeproxy/pipe/components/balance"
	"github.com/pipeproxy/pipe/components/common/register"
	"github.com/pipeproxy/pipe/internal/gcd"
)

var (
	ErrNotHandler = fmt.Errorf("error not handler")
)

const (
	name = "lb"
)

func init() {
	register.Register(name, NewLBWithConfig)
}

type Weight struct {
	Weight  uint `json:",omitempty"`
	Handler http.Handler
}

type Config struct {
	Policy   balance.Policy
	Handlers []*Weight
}

func NewLBWithConfig(conf *Config) (http.Handler, error) {
	switch len(conf.Handlers) {
	case 0:
		return nil, ErrNotHandler
	case 1:
		return conf.Handlers[0].Handler, nil
	}

	var sum uint
	list := make([]uint, 0, len(conf.Handlers))
	for _, weighted := range conf.Handlers {
		if weighted.Weight > 0 {
			list = append(list, weighted.Weight)
			sum += weighted.Weight
		}
	}

	var handlers []http.Handler
	if sum == 0 {
		handlers = make([]http.Handler, 0, len(conf.Handlers))
		for _, weighted := range conf.Handlers {
			handlers = append(handlers, weighted.Handler)
		}
	} else {
		g := gcd.GcdSlice(list)
		handlers = make([]http.Handler, 0, sum/g)
		for _, weighted := range conf.Handlers {
			if weighted.Weight > 0 {
				size := weighted.Weight / g
				for i := uint(0); i != size; i++ {
					handlers = append(handlers, weighted.Handler)
				}
			}
		}
	}
	return NewLB(conf.Policy, handlers), nil
}
