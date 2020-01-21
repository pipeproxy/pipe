package pprof

import (
	"net/http"

	"github.com/wzshiming/pipe/configure"
)

const name = "pprof"

func init() {
	configure.Register(name, NewPprofWithConfig)
}

func NewPprofWithConfig() http.Handler {
	return NewPprof()
}
