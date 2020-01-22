package multi

import (
	"github.com/wzshiming/pipe/configure"
	"github.com/wzshiming/pipe/listener"
)

const name = "multi"

func init() {
	configure.Register(name, NewMultiWithConfig)
}

type Config struct {
	Multi []listener.ListenConfig
}

func NewMultiWithConfig(conf *Config) (listener.ListenConfig, error) {
	if len(conf.Multi) == 0 {
		return nil, ErrNotListener
	}
	return NewMulti(conf.Multi), nil
}
