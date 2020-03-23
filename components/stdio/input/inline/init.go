package inline

import (
	"strings"

	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/components/stdio/input"
)

const name = "inline"

func init() {
	register.Register(name, NewInlineWithConfig)
}

type Config struct {
	Data string
}

func NewInlineWithConfig(conf *Config) input.Input {
	return strings.NewReader(conf.Data)
}
