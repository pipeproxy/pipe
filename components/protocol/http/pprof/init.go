package pprof

import (
	"net/http"

	"github.com/wzshiming/pipe/components/common/register"
)

const (
	name = "pprof"
)

func init() {
	register.Register(name, NewPprofWithConfig)
}

func NewPprofWithConfig() http.Handler {
	return NewPprof()
}
