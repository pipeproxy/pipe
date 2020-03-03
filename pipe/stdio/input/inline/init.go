package inline

import (
	"io/ioutil"
	"strings"

	"github.com/wzshiming/pipe/configure/decode"
	"github.com/wzshiming/pipe/pipe/stdio/input"
)

const name = "inline"

func init() {
	decode.Register(name, NewInlineWithConfig)
}

type Config struct {
	Data string
}

func NewInlineWithConfig(conf *Config) input.Input {
	return ioutil.NopCloser(strings.NewReader(conf.Data))
}
