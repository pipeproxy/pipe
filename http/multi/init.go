package multi

import (
	"net/http"

	"github.com/wzshiming/pipe/configure/manager"
)

const name = "multi"

func init() {
	manager.Register(name, NewMultiWithConfig)
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
