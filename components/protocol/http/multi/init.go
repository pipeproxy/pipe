package multi

import (
	"net/http"

	"github.com/pipeproxy/pipe/components/common/register"
)

const (
	name = "multi"
)

func init() {
	register.Register(name, NewMultiWithConfig)
}

type Config struct {
	Multi []http.Handler
}

func NewMultiWithConfig(conf *Config) (http.Handler, error) {
	switch len(conf.Multi) {
	case 0:
		return nil, ErrNotHandler
	case 1:
		return conf.Multi[0], nil
	}
	return NewMulti(conf.Multi), nil
}
