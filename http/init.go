package http

import (
	"crypto/tls"
	"net/http"

	"github.com/wzshiming/pipe/configure"
	"github.com/wzshiming/pipe/stream"
)

func init() {
	configure.Register("http", NewServerWithConfig)
}

type Config struct {
	Handler http.Handler
	TLS     *tls.Config
}

func NewServerWithConfig(conf *Config) stream.Handler {
	return NewServer(conf.Handler, conf.TLS)
}
