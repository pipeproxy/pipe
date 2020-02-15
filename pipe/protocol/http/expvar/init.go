package expvar

import (
	"expvar"

	"github.com/wzshiming/pipe/configure/decode"
)

const name = "expvar"

func init() {
	decode.Register(name, expvar.Handler)
}
