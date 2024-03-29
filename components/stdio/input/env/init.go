package env

import (
	"bytes"
	"os"

	"github.com/pipeproxy/pipe/components/common/register"
	"github.com/pipeproxy/pipe/components/stdio/input"
)

const (
	name = "env"
)

func init() {
	register.Register(name, NewEnvWithConfig)
}

type Config struct {
	Name string
}

func NewEnvWithConfig(conf *Config) input.Input {
	return input.NewLazyReader(func() (input.Input, error) {
		value, _ := os.LookupEnv(conf.Name)
		return bytes.NewBufferString(value), nil
	})
}
