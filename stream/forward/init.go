package forward

import (
	"context"

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
func NewForwardWithConfig(ctx context.Context, conf *Config) (stream.Handler, error) {
	mux := NewForward(conf.Network, conf.Address)
	return mux, nil
}
