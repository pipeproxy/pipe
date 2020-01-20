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
func NewForwardWithConfig(ctx context.Context, config []byte) (stream.Handler, error) {
	var conf Config
	err := configure.Decode(ctx, config, &conf)
	if err != nil {
		return nil, err
	}
	mux := NewForward(conf.Network, conf.Address)
	return mux, nil
}
