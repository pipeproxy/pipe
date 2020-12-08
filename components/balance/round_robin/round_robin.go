package round_robin

import (
	"math/rand"
	"sync/atomic"

	"github.com/pipeproxy/pipe/components/balance"
)

type RoundRobin struct {
	index uint64
	size  uint64
}

func NewRoundRobin() balance.Policy {
	return &RoundRobin{
		index: rand.Uint64() % 100,
	}
}

func (r *RoundRobin) Init(size uint64) {
	r.size = size
}

func (r *RoundRobin) InUse(fun func(i uint64)) {
	i := (atomic.AddUint64(&r.index, 1) - 1) % r.size
	fun(i)
}

func (r *RoundRobin) Clone() balance.Policy {
	c := NewRoundRobin()
	c.Init(r.size)
	return c
}
