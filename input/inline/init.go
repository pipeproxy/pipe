package file

import (
	"io/ioutil"
	"strings"

	"github.com/wzshiming/pipe/configure/manager"
	"github.com/wzshiming/pipe/input"
)

const name = "inline"

func init() {
	manager.Register(name, NewInlineWithConfig)
}

type Config struct {
	Data string
}

func NewInlineWithConfig(conf *Config) input.Input {
	return ioutil.NopCloser(strings.NewReader(conf.Data))
}
