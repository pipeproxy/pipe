package ref

import (
	"context"

	"github.com/wzshiming/pipe/components/common/register"
)

const (
	name = "ref"
)

func Register(i interface{}) error {
	return register.RegisterWithBuildFunc(name, RefWithConfig, i)
}

type Config struct {
	Name string
}

func RefWithConfig(ctx context.Context, conf *Config, i interface{}) error {
	return Ref(ctx, conf.Name, i)
}
