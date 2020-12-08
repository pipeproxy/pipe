package none

import (
	"github.com/pipeproxy/pipe/components/balance"
)

type None struct{}

func NewNone() balance.Policy {
	return None{}
}

func (None) Init(size uint64) {

}

func (None) InUse(fun func(i uint64)) {
	fun(0)
}

func (None) Clone() balance.Policy {
	return None{}
}
