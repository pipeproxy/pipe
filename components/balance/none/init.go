package none

import (
	"github.com/pipeproxy/pipe/components/common/register"
)

const (
	name = "none"
)

func init() {
	register.Register(name, NewNone)
}
