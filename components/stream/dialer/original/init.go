package original

import (
	"github.com/pipeproxy/pipe/components/common/register"
)

const (
	name = "original"
)

func init() {
	register.Register(name, NewOriginal)
}
