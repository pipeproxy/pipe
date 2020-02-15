package message

import (
	"github.com/wzshiming/pipe/configure/decode"
	"github.com/wzshiming/pipe/pipe/once"
)

const name = "message"

func init() {
	decode.Register(name, NewMessageWithConfig)
}

type Config struct {
	Message string
}

func NewMessageWithConfig(conf *Config) once.Once {
	return Message(conf.Message)
}
