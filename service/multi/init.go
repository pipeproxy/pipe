package multi

import (
	"context"

	"github.com/wzshiming/pipe/configure"
	"github.com/wzshiming/pipe/service"
)

const name = "multi"

func init() {
	configure.Register(name, NewMultiWithConfig)
}

type Config struct {
	Services []service.Service
}

func NewMultiWithConfig(ctx context.Context, config []byte) (service.Service, error) {
	var conf Config
	err := configure.Decode(ctx, config, &conf)
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
