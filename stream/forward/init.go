package forward

import (
	"github.com/wzshiming/pipe/configure"
	"github.com/wzshiming/pipe/stream"
)

const name = "forward"

func init() {
	configure.Register(name, NewForwardWithConfig)
}

type Config struct {
	Network string
	Address string
}

// NewForwardWithConfig create a new forward with config.
func NewForwardWithConfig(conf *Config) stream.Handler {
	return NewForward(conf.Network, conf.Address)
}
