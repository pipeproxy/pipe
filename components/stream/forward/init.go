package forward

import (
	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/components/stream"
	"github.com/wzshiming/pipe/components/stream/dialer"
)

const (
	name = "forward"
)

func init() {
	register.Register(name, NewForwardWithConfig)
}

type Config struct {
	Dialer dialer.Dialer
}

// NewForwardWithConfig create a new forward with config.
func NewForwardWithConfig(conf *Config) stream.Handler {
	return NewForward(conf.Dialer)
}
