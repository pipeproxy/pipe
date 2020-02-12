package message

import (
	"github.com/wzshiming/pipe/configure/manager"
	"github.com/wzshiming/pipe/once"
)

const name = "message"

func init() {
	manager.Register(name, NewMessageWithConfig)
}

type Config struct {
	Message string
}

func NewMessageWithConfig(conf *Config) once.Once {
	return Message(conf.Message)
}
