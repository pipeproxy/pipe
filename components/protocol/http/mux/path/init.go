package path

import (
	"fmt"
	"net/http"

	"github.com/pipeproxy/pipe/components/common/register"
)

const (
	name = "path"
)

func init() {
	register.Register(name, NewPathWithConfig)
}

var (
	ErrNotHandler = fmt.Errorf("error not handler")
	ErrNotRouter  = fmt.Errorf("error not router")
)

type Route struct {
	Prefix  string `json:",omitempty"`
	Path    string `json:",omitempty"`
	Regexp  string `json:",omitempty"`
	Handler http.Handler
}

type Config struct {
	Routes   []*Route
	NotFound http.Handler `json:",omitempty"`
}

func NewPathWithConfig(conf *Config) (http.Handler, error) {
	mux := NewPath()
	mux.NotFound(conf.NotFound)
	for _, route := range conf.Routes {
		if route.Handler == nil {
			return nil, ErrNotHandler
		}
		if route.Path != "" {
			mux.HandlePath(route.Path, route.Handler)
		} else if route.Regexp == "" && route.Prefix != "" {
			err := mux.HandlePrefix(route.Prefix, route.Handler)
			if err != nil {
				return nil, err
			}
		} else if route.Regexp != "" {
			err := mux.HandlePrefixAndRegexp(route.Prefix, route.Regexp, route.Handler)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, ErrNotRouter
		}
	}
	return mux, nil
}
