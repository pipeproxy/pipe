package weighted

import (
	"fmt"
	"net/http"

	"github.com/wzshiming/pipe/configure/manager"
	"github.com/wzshiming/pipe/internal/gcd"
	"github.com/wzshiming/pipe/pipe/protocol/http/poller"
)

var (
	ErrNotWeighted = fmt.Errorf("error not weighted")
)

const name = "weighted"

func init() {
	manager.Register(name, NewWeightedWithConfig)
}

type Weighted struct {
	Weight  uint
	Handler http.Handler
}

type Config struct {
	Weighted []*Weighted
}

func NewWeightedWithConfig(conf *Config) (http.Handler, error) {
	switch len(conf.Weighted) {
	case 0:
		return nil, poller.ErrNotHandler
	case 1:
		return conf.Weighted[0].Handler, nil
	}

	var sum uint
	list := make([]uint, 0, len(conf.Weighted))
	for _, weighted := range conf.Weighted {
		if weighted.Weight > 0 {
			list = append(list, weighted.Weight)
			sum += weighted.Weight
		}
	}

	if sum == 0 {
		return nil, ErrNotWeighted
	}

	g := gcd.GcdSlice(list)

	handlers := make([]http.Handler, 0, sum/g)
	for _, weighted := range conf.Weighted {
		if weighted.Weight > 0 {
			size := weighted.Weight / g
			for i := uint(0); i != size; i++ {
				handlers = append(handlers, weighted.Handler)
			}
		}

	}
	return poller.NewRoundRobin(handlers), nil
}
