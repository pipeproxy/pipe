package h2c

import (
	"net/http"

	"github.com/wzshiming/pipe/configure/decode"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

const name = "h2c"

func init() {
	decode.Register(name, NewH2cWithConfig)
}

type Config struct {
	Handler http.Handler
}

func NewH2cWithConfig(conf *Config) http.Handler {
	return h2c.NewHandler(conf.Handler, &http2.Server{})
}
