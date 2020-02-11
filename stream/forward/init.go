package forward

import (
	"github.com/wzshiming/pipe/configure/manager"
	"github.com/wzshiming/pipe/dialer"
	"github.com/wzshiming/pipe/stream"
)

const name = "forward"

func init() {
	manager.Register(name, NewForwardWithConfig)
}

type Config struct {
	Dialer dialer.Dialer
}

// NewForwardWithConfig create a new forward with config.
func NewForwardWithConfig(conf *Config) stream.Handler {
	return NewForward(conf.Dialer)
}
