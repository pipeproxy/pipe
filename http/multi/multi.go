package multi

import (
	"fmt"
	"net/http"
)

var (
	ErrNotHandler = fmt.Errorf("not handler")
)

type Multi struct {
	multi []http.Handler
}

func NewMulti(multi []http.Handler) *Multi {
	return &Multi{
		multi: multi,
	}
}

func (m *Multi) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	switch len(m.multi) {
	case 0:
	case 1:
		handler := m.multi[0]
		handler.ServeHTTP(rw, r)
	default:
		handlers := m.multi
		for _, handler := range handlers {
			handler.ServeHTTP(rw, r)
		}
	}
}
