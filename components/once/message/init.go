package message

import (
	"github.com/pipeproxy/pipe/components/common/register"
	"github.com/pipeproxy/pipe/components/once"
)

const (
	name = "message"
)

func init() {
	register.Register(name, NewMessageWithConfig)
}

type Config struct {
	Message string
}

func NewMessageWithConfig(conf *Config) once.Once {
	return Message(conf.Message)
}
