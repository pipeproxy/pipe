package http

import (
	"fmt"
	"net/http"

	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/components/stream"
	"github.com/wzshiming/pipe/components/tls"
)

var (
	ErrNotHandler = fmt.Errorf("not handler")
)

func init() {
	register.Register("http", NewServerWithConfig)
}

type Config struct {
	Handler http.Handler
	TLS     tls.TLS
}

func NewServerWithConfig(conf *Config) (stream.Handler, error) {
	if conf.Handler == nil {
		return nil, ErrNotHandler
	}
	if conf.TLS == nil {
		return NewServer(conf.Handler, nil), nil
	}
	return NewServer(conf.Handler, conf.TLS.TLS().Clone()), nil
}
