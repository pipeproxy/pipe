package method

import (
	"net/http"

	"github.com/pipeproxy/pipe/components/common/register"
)

const (
	name = "method"
)

func init() {
	register.Register(name, NewMethodWithConfig)
}

type MethodEnum string

const (
	MethodGet     MethodEnum = "GET"
	MethodHead    MethodEnum = "HEAD"
	MethodPost    MethodEnum = "POST"
	MethodPut     MethodEnum = "PUT"
	MethodPatch   MethodEnum = "PATCH"
	MethodDelete  MethodEnum = "DELETE"
	MethodConnect MethodEnum = "CONNECT"
	MethodOptions MethodEnum = "OPTIONS"
	MethodTrace   MethodEnum = "TRACE"
)

type Route struct {
	Method  MethodEnum
	Handler http.Handler
}

type Config struct {
	Methods  []*Route
	NotFound http.Handler
}

func NewMethodWithConfig(conf *Config) (http.Handler, error) {
	mux := NewMethod()
	mux.NotFound(conf.NotFound)
	for _, route := range conf.Methods {
		if route.Method == "" || route.Handler == nil {
			continue
		}
		mux.Handle(string(route.Method), route.Handler)
	}
	return mux, nil
}
