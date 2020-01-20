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
	Handlers []stream.Handler
}

func NewMultiWithConfig(conf *Config) (stream.Handler, error) {
	switch len(conf.Handlers) {
	case 1:
		return conf.Handlers[0], nil
	case 0:
		return nil, ErrNotHandler
	}
	return NewMulti(conf.Handlers), nil
}
