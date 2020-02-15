package file

import (
	"github.com/wzshiming/pipe/configure/decode"
	"github.com/wzshiming/pipe/pipe/stdio/output"
)

const name = "file"

func init() {
	decode.Register(name, NewFileWithConfig)
}

type Config struct {
	Path string
}

func NewFileWithConfig(conf *Config) (output.Output, error) {
	return NewFile(conf.Path)
}
