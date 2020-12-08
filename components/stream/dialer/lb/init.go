package lb

import (
	"fmt"

	"github.com/pipeproxy/pipe/components/balance"
	"github.com/pipeproxy/pipe/components/common/register"
	"github.com/pipeproxy/pipe/components/stream"
	"github.com/pipeproxy/pipe/internal/gcd"
)

var (
	ErrNotDialer = fmt.Errorf("error not dialer")
)

const (
	name = "lb"
)

func init() {
	register.Register(name, NewLBWithConfig)
}

type Weight struct {
	Weight uint `json:",omitempty"`
	Dialer stream.Dialer
}

type Config struct {
	Policy  balance.Policy
	Dialers []*Weight
}

func NewLBWithConfig(conf *Config) (stream.Dialer, error) {
	switch len(conf.Dialers) {
	case 0:
		return nil, ErrNotDialer
	case 1:
		return conf.Dialers[0].Dialer, nil
	}

	var sum uint
	list := make([]uint, 0, len(conf.Dialers))
	for _, weighted := range conf.Dialers {
		if weighted.Weight > 0 {
			list = append(list, weighted.Weight)
			sum += weighted.Weight
		}
	}

	var dialers []stream.Dialer
	if sum == 0 {
		dialers = make([]stream.Dialer, 0, len(conf.Dialers))
		for _, weighted := range conf.Dialers {
			dialers = append(dialers, weighted.Dialer)
		}
	} else {
		g := gcd.GcdSlice(list)
		dialers = make([]stream.Dialer, 0, sum/g)
		for _, weighted := range conf.Dialers {
			if weighted.Weight > 0 {
				size := weighted.Weight / g
				for i := uint(0); i != size; i++ {
					dialers = append(dialers, weighted.Dialer)
				}
			}
		}
	}
	return NewLB(conf.Policy, dialers), nil
}
