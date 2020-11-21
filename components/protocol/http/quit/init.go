package quit

import (
	"github.com/pipeproxy/pipe/components/common/register"
)

const (
	name = "quit"
)

func init() {
	register.Register(name, NewQuit)
}
