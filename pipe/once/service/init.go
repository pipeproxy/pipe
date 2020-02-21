package service

import (
	"github.com/wzshiming/pipe/configure/decode"
	"github.com/wzshiming/pipe/pipe/once"
	"github.com/wzshiming/pipe/pipe/service"
)

const name = "service"

func init() {
	decode.Register(name, NewServiceWithConfig)
}

type Config struct {
	Service service.Service
}

func NewServiceWithConfig(conf *Config) once.Once {
	return NewService(conf.Service)
}
