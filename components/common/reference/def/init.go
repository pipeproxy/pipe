package def

import (
	"context"

	"github.com/wzshiming/funcfg/define"
	"github.com/wzshiming/pipe/components/common/register"
)

const (
	name = "def"
)

func Register(i interface{}) error {
	return register.RegisterWithBuildFunc(name, DefWithConfig, i)
}

type Config struct {
	Name string
	Def  define.Self
}

func DefWithConfig(ctx context.Context, conf *Config, i interface{}) error {
	return Def(ctx, conf.Name, conf.Def, i)
}
