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
	handlers := m.multi
	for _, handler := range handlers {
		handler.ServeHTTP(rw, r)
	}
}
