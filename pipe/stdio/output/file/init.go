package file

import (
	"github.com/wzshiming/pipe/configure/manager"
	"github.com/wzshiming/pipe/pipe/stdio/output"
)

const name = "file"

func init() {
	manager.Register(name, NewFileWithConfig)
}

type Config struct {
	Path string
}

func NewFileWithConfig(conf *Config) (output.Output, error) {
	return NewFile(conf.Path)
}
