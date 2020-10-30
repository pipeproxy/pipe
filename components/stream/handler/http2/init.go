package http

import (
	"fmt"
	"net/http"

	"github.com/pipeproxy/pipe/components/common/register"
	"github.com/pipeproxy/pipe/components/stream"
	"github.com/pipeproxy/pipe/components/tls"
)

var (
	ErrNotHandler = fmt.Errorf("not handler")
)

func init() {
	register.Register("http2", NewServerWithConfig)
}

type Config struct {
	Handler http.Handler
	TLS     tls.TLS `json:",omitempty"`
}

func NewServerWithConfig(conf *Config) (stream.Handler, error) {
	if conf.Handler == nil {
		return nil, ErrNotHandler
	}
	return NewServer(conf.Handler, conf.TLS), nil
}
