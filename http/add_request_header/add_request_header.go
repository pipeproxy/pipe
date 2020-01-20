package add_request_header

import (
	"net/http"
)

type AddRequestHeader struct {
	key    string
	values []string
}

func NewAddRequestHeader(key string, values []string) *AddRequestHeader {
	return &AddRequestHeader{key, values}
}

func (a *AddRequestHeader) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	r.Header[a.key] = append(r.Header[a.key], a.values...)
}
