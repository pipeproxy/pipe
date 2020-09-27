package lb

import (
	"fmt"

	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/components/stream"
	"github.com/wzshiming/pipe/internal/gcd"
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

type LoadBalancePolicyEnum string

const (
	EnumRoundRobin LoadBalancePolicyEnum = "round_robin"
	EnumRandom     LoadBalancePolicyEnum = "random"
)

type Weight struct {
	Weight  uint `json:",omitempty"`
	Handler stream.Handler
}

type Config struct {
	Policy   LoadBalancePolicyEnum `json:",omitempty"`
	Handlers []*Weight
}

func NewLBWithConfig(conf *Config) (stream.Handler, error) {
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

	var dialers []stream.Handler
	if sum == 0 {
		dialers = make([]stream.Handler, 0, len(conf.Handlers))
		for _, weighted := range conf.Handlers {
			dialers = append(dialers, weighted.Handler)
		}
	} else {
		g := gcd.GcdSlice(list)
		dialers = make([]stream.Handler, 0, sum/g)
		for _, weighted := range conf.Handlers {
			if weighted.Weight > 0 {
				size := weighted.Weight / g
				for i := uint(0); i != size; i++ {
					dialers = append(dialers, weighted.Handler)
				}
			}
		}
	}

	switch conf.Policy {
	case EnumRandom:
		return NewRandom(dialers), nil
	default: // EnumRoundRobin
		return NewRoundRobin(dialers), nil
	}
}
