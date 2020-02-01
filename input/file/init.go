package file

import (
	"bytes"
	"io/ioutil"

	"github.com/wzshiming/pipe/configure"
	"github.com/wzshiming/pipe/input"
)

const name = "file"

func init() {
	configure.Register(name, NewFileWithConfig)
}

type Config struct {
	Path string
}

func NewFileWithConfig(conf *Config) (input.Input, error) {
	data, err := ioutil.ReadFile(conf.Path)
	if err != nil {
		return nil, err
	}
	return ioutil.NopCloser(bytes.NewReader(data)), nil
}
