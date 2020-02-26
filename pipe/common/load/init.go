package load

import (
	"context"
	"io"

	"github.com/wzshiming/pipe/configure/decode"
)

const name = "load"

func Register(i interface{}) error {
	return decode.RegisterWithBuildFunc(name, LoadWithConfig, i)
}

type Config struct {
	Load io.ReadCloser
}

func LoadWithConfig(ctx context.Context, conf *Config, i interface{}) error {
	return Load(ctx, conf.Load, i)
}
