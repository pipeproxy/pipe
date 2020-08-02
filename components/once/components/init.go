package components

import (
	"github.com/wzshiming/funcfg/define"
	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/components/once"
)

const (
	name = "components"
)

func init() {
	register.Register(name, newComponentsWithConfig)
}

type Config struct {
	Components []define.Any
}

func newComponentsWithConfig(conf *Config) once.Once {
	return none{}
}
