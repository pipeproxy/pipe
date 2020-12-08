package round_robin

import (
	"sync/atomic"

	"github.com/pipeproxy/pipe/components/balance"
)

type RoundRobin struct {
	index uint64
}

func NewRoundRobin() balance.Policy {
	return &RoundRobin{}
}

func (r *RoundRobin) InUse(size uint64, fun func(i uint64)) {
	i := (atomic.AddUint64(&r.index, 1) - 1) % size
	fun(i)
}

func (RoundRobin) Policy() balance.PolicyEnum {
	return balance.EnumPolicyRoundRobin
}
