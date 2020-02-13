package pprof

import (
	"net/http"

	"github.com/wzshiming/pipe/configure/manager"
)

const name = "pprof"

func init() {
	manager.Register(name, NewPprofWithConfig)
}

func NewPprofWithConfig() http.Handler {
	return NewPprof()
}
