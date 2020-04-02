package sample

import (
	"github.com/wzshiming/funcfg/define"
	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/components/once"
	"github.com/wzshiming/pipe/components/once/multi"
	osvc "github.com/wzshiming/pipe/components/once/service"
	"github.com/wzshiming/pipe/components/service"
)

const name = "sample"

func init() {
	register.Register(name, NewConfigWithConfig)
}

type Config struct {
	Components []define.Any
	Pipe       service.Service
	Init       []once.Once
}

func NewConfigWithConfig(conf *Config) once.Once {
	return multi.NewMulti(append(conf.Init, osvc.NewService(conf.Pipe)))
}
