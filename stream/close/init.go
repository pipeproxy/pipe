package close

import (
	"github.com/wzshiming/pipe/configure"
	"github.com/wzshiming/pipe/stream"
)

const name = "close"

func init() {
	configure.Register(name, NewCloseWithConfig)
}

type Config struct {
	Network string
	Address string
}

// NewCloseWithConfig create a new close with config.
func NewCloseWithConfig() stream.Handler {
	return NewClose()
}
