package expvar

import (
	"expvar"

	"github.com/wzshiming/pipe/components/common/register"
)

const name = "expvar"

func init() {
	register.Register(name, expvar.Handler)
}
