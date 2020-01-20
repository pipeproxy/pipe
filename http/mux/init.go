package mux

import (
	"fmt"
	"net/http"

	"github.com/wzshiming/pipe/configure"
)

const name = "mux"

func init() {
	configure.Register(name, NewMuxWithConfig)
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
			mux.HandlePrefix(route.Path, route.Handler)
		}
	}
	return mux, nil
}
