package remove_response_header

import (
	"net/http"
)

type RemoveResponseHeader struct {
	key string
}

func NewRemoveResponseHeader(key string) *RemoveResponseHeader {
	return &RemoveResponseHeader{key}
}

func (e *RemoveResponseHeader) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	header := rw.Header()
	delete(header, e.key)
}
