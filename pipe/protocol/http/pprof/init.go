package pprof

import (
	"net/http"

	"github.com/wzshiming/pipe/configure/decode"
)

const name = "pprof"

func init() {
	decode.Register(name, NewPprofWithConfig)
}

func NewPprofWithConfig() http.Handler {
	return NewPprof()
}
