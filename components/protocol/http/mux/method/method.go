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
	methods  map[string]http.Handler
	notFound http.Handler
}

func NewMethod() *Method {
	p := &Method{
		methods: map[string]http.Handler{},
	}
	return p
}

func (h *Method) NotFound(handler http.Handler) {
	h.notFound = handler
}

func (h *Method) Handle(method string, handler http.Handler) {
	h.methods[method] = handler
}

// Handler returns methods route handler.
func (h *Method) Handler(method string) (handler http.Handler, err error) {
	handler, ok := h.methods[method]
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
