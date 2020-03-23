package load

import (
	"context"
	"io"

	"github.com/wzshiming/pipe/components/common/register"
)

const name = "load"

func Register(i interface{}) error {
	return register.RegisterWithBuildFunc(name, LoadWithConfig, i)
}

type Config struct {
	Load io.Reader
}

func LoadWithConfig(ctx context.Context, conf *Config, i interface{}) error {
	return Load(ctx, conf.Load, i)
}
