package none

import (
	"github.com/wzshiming/funcfg/define"
	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/components/once"
)

const (
	name = "none"
)

func init() {
	register.Register(name, NewNoneWithConfig)
}

type Config struct {
	Any define.Any
}

func NewNoneWithConfig(conf *Config) (once.Once, error) {
	return None{}, nil
}
