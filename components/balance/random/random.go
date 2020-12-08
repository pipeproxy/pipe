package random

import (
	"math/rand"

	"github.com/pipeproxy/pipe/components/balance"
)

type Random struct{}

func NewRandom() balance.Policy {
	return Random{}
}

func (Random) InUse(size uint64, fun func(i uint64)) {
	i := rand.Uint64() % size
	fun(i)
}

func (Random) Policy() balance.PolicyEnum {
	return balance.EnumPolicyRandom
}
