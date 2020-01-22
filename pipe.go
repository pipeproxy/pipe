package pipe

import (
	"context"

	"github.com/wzshiming/pipe/configure"
	"github.com/wzshiming/pipe/service"
)

func NewPipeWithConfig(ctx context.Context, config []byte) (service.Service, error) {
	var conf Config
	err := configure.Decode(ctx, config, &conf)
	if err != nil {
		return nil, err
	}
	return conf.Pipe, nil
}
