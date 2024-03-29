package inline

import (
	"bytes"

	"github.com/pipeproxy/pipe/components/common/register"
	"github.com/pipeproxy/pipe/components/stdio/input"
)

const (
	name = "inline"
)

func init() {
	register.Register(name, NewInlineWithConfig)
}

type Config struct {
	Data string
}

func NewInlineWithConfig(conf *Config) input.Input {
	return bytes.NewBufferString(conf.Data)
}
