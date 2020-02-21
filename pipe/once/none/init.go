package none

import (
	"github.com/wzshiming/pipe/configure/decode"
	"github.com/wzshiming/pipe/pipe/once"
)

const name = "none"

func init() {
	decode.Register(name, NewMultiWithConfig)
}

type Config struct {
	Any interface{}
}

func NewMultiWithConfig(conf *Config) (once.Once, error) {
	return None{}, nil
}
