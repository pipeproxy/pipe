package http

import (
	"fmt"
	"net/http"

	"github.com/wzshiming/pipe/configure/manager"
	"github.com/wzshiming/pipe/stream"
	"github.com/wzshiming/pipe/tls"
)

var (
	ErrNotHandler = fmt.Errorf("not handler")
)

func init() {
	manager.Register("http", NewServerWithConfig)
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
