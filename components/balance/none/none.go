package none

import (
	"github.com/pipeproxy/pipe/components/balance"
)

type None struct{}

func NewNone() balance.Policy {
	return None{}
}

func (None) InUse(size uint64, fun func(i uint64)) {
	fun(0)
}

func (None) Policy() balance.PolicyEnum {
	return balance.EnumPolicyNone
}
