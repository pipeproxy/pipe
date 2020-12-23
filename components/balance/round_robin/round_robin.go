package round_robin

import (
	"sync/atomic"

	"github.com/pipeproxy/pipe/components/balance"
	"github.com/pipeproxy/pipe/internal/rand"
)

type RoundRobin struct {
	index uint64
	size  uint64
}

func NewRoundRobin() balance.Policy {
	return &RoundRobin{}
}

func (r *RoundRobin) Init(size uint64) {
	r.size = size
	r.index = rand.Uint64() % size
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
