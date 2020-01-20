package add_request_header

import (
	"net/http"

	"github.com/wzshiming/pipe/http/template"
)

type AddRequestHeader struct {
	key   string
	value template.Format
}

func NewAddRequestHeader(key string, value template.Format) *AddRequestHeader {
	return &AddRequestHeader{key, value}
}

func (a *AddRequestHeader) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	r.Header[a.key] = append(r.Header[a.key], a.value.FormatString(r))
}
