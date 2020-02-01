package file

import (
	"io/ioutil"
	"strings"

	"github.com/wzshiming/pipe/configure"
	"github.com/wzshiming/pipe/input"
)

const name = "inline"

func init() {
	configure.Register(name, NewInlineWithConfig)
}

type Config struct {
	Data string
}

func NewInlineWithConfig(conf *Config) input.Input {
	return ioutil.NopCloser(strings.NewReader(conf.Data))
}
