package quit

import (
	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/components/protocol/http"
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
