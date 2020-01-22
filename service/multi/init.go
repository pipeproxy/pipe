package multi

import (
	"github.com/wzshiming/pipe/configure"
	"github.com/wzshiming/pipe/service"
)

const name = "multi"

func init() {
	configure.Register(name, NewMultiWithConfig)
}

type Config struct {
	Multi []service.Service
}

func NewMultiWithConfig(conf *Config) (service.Service, error) {
	if len(conf.Multi) == 0 {
		return nil, ErrNotServer
	}
	return NewMulti(conf.Multi), nil
}
