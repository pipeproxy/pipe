package service

import (
	"github.com/pipeproxy/pipe/components/common/register"
	"github.com/pipeproxy/pipe/components/once"
	"github.com/pipeproxy/pipe/components/service"
)

const (
	name = "service"
)

func init() {
	register.Register(name, NewServiceWithConfig)
}

type Config struct {
	Service service.Service
}

func NewServiceWithConfig(conf *Config) once.Once {
	return NewService(conf.Service)
}
