package file

import (
	"github.com/wzshiming/pipe/configure"
	"github.com/wzshiming/pipe/output"
)

const name = "file"

func init() {
	configure.Register(name, NewFileWithConfig)
}

type Config struct {
	Path string
}

func NewFileWithConfig(conf *Config) (output.Output, error) {
	return NewFile(conf.Path)
}
