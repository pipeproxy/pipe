package method

import (
	"net/http"

	"github.com/pipeproxy/pipe/internal/http/template"
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

func (m *Method) NotFound(handler http.Handler) {
	m.notFound = handler
}

func (m *Method) Handle(method string, handler http.Handler) {
	m.methods[method] = handler
}

// Handler returns methods route handler.
func (m *Method) Handler(method string) (handler http.Handler) {
	handler, ok := m.methods[method]
	if ok {
		return handler
	}
	if m.notFound == nil {
		return template.NotFoundHandler
	}
	return m.notFound
}

func (m *Method) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	m.Handler(r.Host).ServeHTTP(rw, r)
}
