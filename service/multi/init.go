package multi

import (
	"context"

	"github.com/wzshiming/pipe/decode"
	"github.com/wzshiming/pipe/service"
)

const name = "multi"

func init() {
	decode.Register(name, NewMultiWithConfig)
}

type Config struct {
	Services []service.Service
}

func NewMultiWithConfig(ctx context.Context, config []byte) (service.Service, error) {
	var conf Config
	err := decode.Decode(ctx, config, &conf)
	if err != nil {
		return nil, err
	}
	switch len(conf.Services) {
	case 1:
		return conf.Services[0], nil
	case 0:
		return nil, ErrNotServer
	}
	return NewMulti(conf.Services)
}
