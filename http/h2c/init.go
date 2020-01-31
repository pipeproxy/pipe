package h2c

import (
	"net/http"

	"golang.org/x/net/http2"

	"github.com/wzshiming/pipe/configure"
	"golang.org/x/net/http2/h2c"
)

const name = "h2c"

func init() {
	configure.Register(name, NewH2cWithConfig)
}

type Config struct {
	Handler http.Handler
}

func NewH2cWithConfig(conf *Config) http.Handler {
	return h2c.NewHandler(conf.Handler, &http2.Server{})
}
