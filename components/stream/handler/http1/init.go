package http

import (
	"fmt"
	"net/http"

	"github.com/pipeproxy/pipe/components/common/register"
	"github.com/pipeproxy/pipe/components/stream"
)

var (
	ErrNotHandler = fmt.Errorf("not handler")
)

func init() {
	register.Register("http1", NewServerWithConfig)
}

type Config struct {
	Handler http.Handler
}

func NewServerWithConfig(conf *Config) (stream.Handler, error) {
	if conf.Handler == nil {
		return nil, ErrNotHandler
	}
	return NewServer(conf.Handler), nil
}
