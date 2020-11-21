package pprof

import (
	"github.com/pipeproxy/pipe/components/common/register"
)

const (
	name = "pprof"
)

func init() {
	register.Register(name, NewPprof)
}
