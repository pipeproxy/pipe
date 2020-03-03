package host

import (
	"fmt"
	"net/http"

	"github.com/wzshiming/pipe/configure/decode"
)

const name = "host"

func init() {
	decode.Register(name, NewHostWithConfig)
}

var (
	ErrNotHandler = fmt.Errorf("error not handler")
	ErrNotDomain  = fmt.Errorf("error not router")
)

type Route struct {
	Domain  string
	Handler http.Handler
}

type Config struct {
	Hosts    []*Route
	NotFound http.Handler
}

func NewHostWithConfig(conf *Config) (http.Handler, error) {
	mux := NewHost()
	mux.NotFound(conf.NotFound)
	for _, route := range conf.Hosts {
		if route.Handler == nil {
			return nil, ErrNotHandler
		}
		if route.Domain == "" {
			return nil, ErrNotDomain
		}
		mux.Handle(route.Domain, route.Handler)
	}
	return mux, nil
}
