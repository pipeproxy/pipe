package hosts

import (
	"fmt"
	"net/http"

	"github.com/wzshiming/pipe/components/common/register"
)

const (
	name = "hosts"
)

func init() {
	register.Register(name, NewHostsWithConfig)
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

func NewHostsWithConfig(conf *Config) (http.Handler, error) {
	mux := NewHosts()
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
