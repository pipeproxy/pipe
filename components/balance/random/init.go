package random

import (
	"github.com/pipeproxy/pipe/components/common/register"
)

const (
	name = "random"
)

func init() {
	register.Register(name, NewRandom)
}
