package method

import (
	"fmt"
	"net/http"
)

var (
	ErrNotFound = fmt.Errorf("error not found")
)

// Method is an host multiplexer.
type Method struct {
	method   map[string]http.Handler
	notFound http.Handler
}

func NewMethod() *Method {
	p := &Method{
		method: map[string]http.Handler{},
	}
	return p
}

func (h *Method) NotFound(handler http.Handler) error {
	h.notFound = handler
	return nil
}

func (h *Method) Handle(host string, handler http.Handler) error {
	h.method[host] = handler
	return nil
}

// Handler returns method route handler.
func (h *Method) Handler(host string) (handler http.Handler, err error) {
	handler, ok := h.method[host]
	if ok {
		return handler, nil
	}
	if h.notFound == nil {
		return nil, ErrNotFound
	}
	return h.notFound, nil
}

func (h *Method) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	host := r.Host
	handler, err := h.Handler(host)
	if err != nil || handler == nil {
		handler = http.HandlerFunc(http.NotFound)
	}
	handler.ServeHTTP(rw, r)
}
