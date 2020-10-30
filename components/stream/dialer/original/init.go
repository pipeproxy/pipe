package original

import (
	"github.com/wzshiming/pipe/components/common/register"
)

const (
	name = "original"
)

func init() {
	register.Register(name, NewOriginal)
}
