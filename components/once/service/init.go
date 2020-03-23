package service

import (
	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/components/once"
	"github.com/wzshiming/pipe/components/service"
)

const name = "service"

func init() {
	register.Register(name, NewServiceWithConfig)
}

type Config struct {
	Service service.Service
}

func NewServiceWithConfig(conf *Config) once.Once {
	return NewService(conf.Service)
}
