package multi

import (
	"github.com/pipeproxy/pipe/components/common/register"
	"github.com/pipeproxy/pipe/components/once"
)

const (
	name = "multi"
)

func init() {
	register.Register(name, NewMultiWithConfig)
}

type Config struct {
	Multi []once.Once
}

func NewMultiWithConfig(conf *Config) (once.Once, error) {
	switch len(conf.Multi) {
	case 0:
		return nil, ErrNotServer
	case 1:
		return conf.Multi[0], nil
	}
	return NewMulti(conf.Multi), nil
}
