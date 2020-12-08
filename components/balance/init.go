package balance

import (
	"github.com/pipeproxy/pipe/components/common/types"
)

func init() {
	var policy Policy
	types.Register(&policy)
}

type Policy interface {
	Init(size uint64)
	InUse(fun func(i uint64))
	Clone() Policy
}
