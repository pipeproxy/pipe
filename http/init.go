package http

import (
	"net/http"

	"github.com/wzshiming/pipe/configure"
	"github.com/wzshiming/pipe/stream"
	"github.com/wzshiming/pipe/tls"
)

func init() {
	configure.Register("http", NewServerWithConfig)
}

type Config struct {
	Handler http.Handler
	TLS     tls.TLS
}

func NewServerWithConfig(conf *Config) stream.Handler {
	if conf.TLS == nil {
		return NewServer(conf.Handler, nil)
	}
	return NewServer(conf.Handler, conf.TLS.TLS())
}
