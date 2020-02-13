package multi

import (
	"github.com/wzshiming/pipe/configure/manager"
	"github.com/wzshiming/pipe/pipe/service"
)

const name = "multi"

func init() {
	manager.Register(name, NewMultiWithConfig)
}

type Config struct {
	Multi []service.Service
}

func NewMultiWithConfig(conf *Config) (service.Service, error) {
	switch len(conf.Multi) {
	case 0:
		return nil, ErrNotServer
	case 1:
		return conf.Multi[0], nil
	}
	return NewMulti(conf.Multi), nil
}
