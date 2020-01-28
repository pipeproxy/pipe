package multi

import (
	"github.com/wzshiming/pipe/configure"
	"github.com/wzshiming/pipe/stream"
)

const name = "multi"

func init() {
	configure.Register(name, NewMultiWithConfig)
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
