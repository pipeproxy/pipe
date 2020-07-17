package file

import (
	"bytes"
	"io/ioutil"

	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/components/stdio/input"
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

func NewFileWithConfig(conf *Config) (input.Input, error) {
	data, err := ioutil.ReadFile(conf.Path)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(data), nil
}
