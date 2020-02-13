package multi

import (
	"github.com/wzshiming/pipe/configure/manager"
	"github.com/wzshiming/pipe/pipe/stream/listener"
)

const name = "multi"

func init() {
	manager.Register(name, NewMultiWithConfig)
}

type Config struct {
	Multi []listener.ListenConfig
}

func NewMultiWithConfig(conf *Config) (listener.ListenConfig, error) {
	switch len(conf.Multi) {
	case 0:
		return nil, ErrNotListener
	case 1:
		return conf.Multi[0], nil
	}
	return NewMulti(conf.Multi), nil
}
