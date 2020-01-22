package ref

import (
	"context"
	"fmt"

	"github.com/wzshiming/pipe/components"
	"github.com/wzshiming/pipe/configure"
	"github.com/wzshiming/pipe/service"
)

var (
	ErrNotService = fmt.Errorf("not service")
)

const name = "ref"

func init() {
	configure.Register(name, NewRefWithConfig)
}

type Config struct {
	Ref string
}

func NewRefWithConfig(ctx context.Context, conf *Config) (service.Service, error) {
	components, ok := components.GetCtxComponents(ctx)
	if !ok || components == nil || components.Services == nil {
		return nil, ErrNotService
	}
	service, ok := components.Services[conf.Ref]
	if !ok {
		return nil, fmt.Errorf("%s: %w", conf.Ref, ErrNotService)
	}
	return service, nil
}
