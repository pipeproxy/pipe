package pipe

import (
	"context"

	"github.com/wzshiming/pipe/service"

	"github.com/wzshiming/pipe/decode"
)

func NewPipeWithConfig(ctx context.Context, config []byte) (service.Service, error) {
	var conf Config
	err := decode.Decode(ctx, config, &conf)
	if err != nil {
		return nil, err
	}
	return conf.Pipe, nil
}
