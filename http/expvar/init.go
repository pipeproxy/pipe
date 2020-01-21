package expvar

import (
	"expvar"

	"github.com/wzshiming/pipe/configure"
)

const name = "expvar"

func init() {
	configure.Register(name, expvar.Handler)
}
