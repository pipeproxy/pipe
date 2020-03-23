package reference

import (
	"context"

	"github.com/wzshiming/funcfg/define"
	"github.com/wzshiming/pipe/components/common/register"
)

const (
	nameDef = "def"
	nameRef = "ref"
)

func Register(i interface{}) error {
	err := register.RegisterWithBuildFunc(nameDef, DefWithConfig, i)
	if err != nil {
		return err
	}
	return register.RegisterWithBuildFunc(nameRef, RefWithConfig, i)
}

type Config struct {
	Name string
	Def  define.Self
}

func DefWithConfig(ctx context.Context, conf *Config, i interface{}) error {
	return Def(ctx, conf.Name, conf.Def, i)
}

func RefWithConfig(ctx context.Context, conf *Config, i interface{}) error {
	return Ref(ctx, conf.Name, conf.Def, i)
}
