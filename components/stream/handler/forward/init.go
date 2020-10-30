package forward

import (
	"github.com/pipeproxy/pipe/components/common/register"
	"github.com/pipeproxy/pipe/components/stream"
)

const (
	name = "forward"
)

func init() {
	register.Register(name, NewForwardWithConfig)
}

type Config struct {
	Dialer stream.Dialer
}

// NewForwardWithConfig create a new forward with config.
func NewForwardWithConfig(conf *Config) stream.Handler {
	return NewForward(conf.Dialer)
}
