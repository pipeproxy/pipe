package multi

import (
	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/components/service"
)

const (
	name = "multi"
)

func init() {
	register.Register(name, NewMultiWithConfig)
}

type Config struct {
	Multi []service.Service
}

func NewMultiWithConfig(conf *Config) (service.Service, error) {
	switch len(conf.Multi) {
	case 0:
		return nil, ErrNotServer
	}
	return NewMulti(conf.Multi), nil
}
