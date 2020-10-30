package file

import (
	"github.com/pipeproxy/pipe/components/common/register"
	"github.com/pipeproxy/pipe/components/stdio/output"
)

const (
	name = "file"
)

func init() {
	register.Register(name, NewFileWithConfig)
}

type Config struct {
	Path string
}

func NewFileWithConfig(conf *Config) (output.Output, error) {
	return NewFile(conf.Path)
}
