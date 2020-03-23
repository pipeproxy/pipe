package multi

import (
	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/components/stream"
)

const name = "multi"

func init() {
	register.Register(name, NewMultiWithConfig)
}

type Config struct {
	Multi []stream.Handler
}

func NewMultiWithConfig(conf *Config) (stream.Handler, error) {
	switch len(conf.Multi) {
	case 0:
		return nil, ErrNotHandler
	case 1:
		return conf.Multi[0], nil
	}
	return NewMulti(conf.Multi), nil
}
