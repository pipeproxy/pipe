package config

import (
	"github.com/wzshiming/pipe/configure/decode"
	"github.com/wzshiming/pipe/pipe/once"
	"github.com/wzshiming/pipe/pipe/once/multi"
	osvc "github.com/wzshiming/pipe/pipe/once/service"
	"github.com/wzshiming/pipe/pipe/service"
)

const name = "config"

func init() {
	decode.Register(name, NewConfigWithConfig)
}

type Config struct {
	Components []interface{}
	Pipe       service.Service
	Init       []once.Once
}

func NewConfigWithConfig(conf *Config) once.Once {
	return multi.NewMulti(append(conf.Init, osvc.NewService(conf.Pipe)))
}
