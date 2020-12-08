package round_robin

import (
	"github.com/pipeproxy/pipe/components/common/register"
)

const (
	name = "round_robin"
)

func init() {
	register.Register(name, NewRoundRobin)
}
