package mux

import (
	"fmt"
	"net/http"

	"github.com/wzshiming/pipe/configure/decode"
)

const name = "mux"

func init() {
	decode.Register(name, NewMuxWithConfig)
}

var (
	ErrNotHandler = fmt.Errorf("error not handler")
)

type Route struct {
	Prefix  string
	Path    string
	Regexp  string
	Handler http.Handler
}

type Config struct {
	Routes   []*Route
	NotFound http.Handler
}

func NewMuxWithConfig(conf *Config) (http.Handler, error) {
	mux := NewMux()
	mux.NotFound(conf.NotFound)

	for _, route := range conf.Routes {
		if route.Handler == nil {
			return nil, ErrNotHandler
		}
		if route.Path != "" {
			mux.HandlePath(route.Path, route.Handler)
		}
		if route.Prefix != "" {
			mux.HandlePrefix(route.Prefix, route.Handler)
		}
	}
	return mux, nil
}