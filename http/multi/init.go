package multi

import (
	"net/http"

	"github.com/wzshiming/pipe/configure"
)

const name = "multi"

func init() {
	configure.Register(name, NewMultiWithConfig)
}

type Config struct {
	Multi []http.Handler
}

func NewMultiWithConfig(conf *Config) (http.Handler, error) {
	switch len(conf.Multi) {
	case 1:
		return conf.Multi[0], nil
	case 0:
		return nil, ErrNotHandler
	}
	return NewMulti(conf.Multi), nil
}
