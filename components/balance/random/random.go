package random

import (
	"github.com/pipeproxy/pipe/components/balance"
	"github.com/pipeproxy/pipe/internal/rand"
)

type Random struct {
	size uint64
}

func NewRandom() balance.Policy {
	return &Random{}
}

func (r *Random) Init(size uint64) {
	r.size = size
}

func (r *Random) InUse(fun func(i uint64)) {
	i := rand.Uint64() % r.size
	fun(i)
}

func (r *Random) Clone() balance.Policy {
	c := NewRandom()
	c.Init(r.size)
	return c
}
