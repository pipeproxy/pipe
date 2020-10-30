package components

import (
	"github.com/pipeproxy/pipe/components/common/register"
	"github.com/pipeproxy/pipe/components/once"
	"github.com/wzshiming/funcfg/define"
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
