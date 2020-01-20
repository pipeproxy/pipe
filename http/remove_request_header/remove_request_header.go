package remove_request_header

import (
	"net/http"
)

type RemoveRequestHeader struct {
	key string
}

func NewRemoveRequestHeader(key string) *RemoveRequestHeader {
	return &RemoveRequestHeader{key}
}

func (a *RemoveRequestHeader) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	delete(r.Header, a.key)
}
