package add_response_header

import (
	"net/http"
)

type AddResponseHeader struct {
	key    string
	values []string
}

func NewAddResponseHeader(key string, values []string) *AddResponseHeader {
	return &AddResponseHeader{key, values}
}

func (a *AddResponseHeader) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	header := rw.Header()
	header[a.key] = append(header[a.key], a.values...)
}
