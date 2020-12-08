package balance

import (
	"github.com/pipeproxy/pipe/components/common/types"
)

func init() {
	var policy Policy
	types.Register(&policy)
}

type Policy interface {
	InUse(size uint64, fun func(i uint64))
	Policy() PolicyEnum
}

type PolicyEnum string

const (
	EnumPolicyNone       PolicyEnum = "none"
	EnumPolicyRandom     PolicyEnum = "random"
	EnumPolicyRoundRobin PolicyEnum = "round_robin"
)
