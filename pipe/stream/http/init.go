package http

import (
	"fmt"
	"net/http"

	"github.com/wzshiming/pipe/configure/decode"
	"github.com/wzshiming/pipe/pipe/stream"
	"github.com/wzshiming/pipe/pipe/tls"
)

var (
	ErrNotHandler = fmt.Errorf("not handler")
)

func init() {
	decode.Register("http", NewServerWithConfig)
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
