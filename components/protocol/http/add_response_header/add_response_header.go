package add_response_header

import (
	"net/http"

	"github.com/wzshiming/pipe/components/protocol/http/template"
)

type AddResponseHeader struct {
	key   string
	value template.Format
}

func NewAddResponseHeader(key string, value template.Format) *AddResponseHeader {
	return &AddResponseHeader{key, value}
}

func (a *AddResponseHeader) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	header := rw.Header()
	header[a.key] = append(header[a.key], a.value.FormatString(r))
}
