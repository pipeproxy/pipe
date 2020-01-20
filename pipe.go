package pipe

import (
	"context"

	"github.com/wzshiming/pipe/components"
	"github.com/wzshiming/pipe/configure"
	"github.com/wzshiming/pipe/service"
)

func NewPipeWithConfig(ctx context.Context, config []byte) (service.Service, error) {
	var cconf configComponents
	err := configure.Decode(ctx, config, &cconf)
	if err != nil {
		return nil, err
	}
	ctx = components.PutCtxComponents(ctx, cconf.Components)
	var conf Config
	err = configure.Decode(ctx, config, &conf)
	if err != nil {
		return nil, err
	}
	return conf.Pipe, nil
}
