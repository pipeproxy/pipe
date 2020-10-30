package quit

import (
	"github.com/pipeproxy/pipe/components/common/register"
	"github.com/pipeproxy/pipe/components/protocol/http"
)

const (
	name = "quit"
)

func init() {
	register.Register(name, NewPprofWithConfig)
}

func NewPprofWithConfig() http.Handler {
	return NewQuit()
}
