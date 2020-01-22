package ref

import (
	"context"
	"fmt"

	"github.com/wzshiming/pipe/listener"

	"github.com/wzshiming/pipe/components"
	"github.com/wzshiming/pipe/configure"
)

var (
	ErrNotListener = fmt.Errorf("not listener")
)

const name = "ref"

func init() {
	configure.Register(name, NewRefWithConfig)
}

type Config struct {
	Ref string
}

func NewRefWithConfig(ctx context.Context, conf *Config) (listener.ListenConfig, error) {
	components, ok := components.GetCtxComponents(ctx)
	if !ok || components == nil || components.Listeners == nil {
		return nil, fmt.Errorf("%s: %w", conf.Ref, ErrNotListener)
	}
	listenConfig, ok := components.Listeners[conf.Ref]
	if !ok {
		return nil, fmt.Errorf("%s: %w", conf.Ref, ErrNotListener)
	}
	return listenConfig, nil
}
