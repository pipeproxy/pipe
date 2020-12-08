package least

import (
	"github.com/pipeproxy/pipe/components/common/register"
)

const (
	name = "least"
)

func init() {
	register.Register(name, NewLeast)
}
