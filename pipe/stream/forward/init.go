package forward

import (
	"github.com/wzshiming/pipe/configure/decode"
	"github.com/wzshiming/pipe/pipe/stream"
	"github.com/wzshiming/pipe/pipe/stream/dialer"
)

const name = "forward"

func init() {
	decode.Register(name, NewForwardWithConfig)
}

type Config struct {
	Dialer dialer.Dialer
}

// NewForwardWithConfig create a new forward with config.
func NewForwardWithConfig(conf *Config) stream.Handler {
	return NewForward(conf.Dialer)
}
