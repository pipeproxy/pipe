package expvar

import (
	"expvar"

	"github.com/wzshiming/pipe/configure/manager"
)

const name = "expvar"

func init() {
	manager.Register(name, expvar.Handler)
}
