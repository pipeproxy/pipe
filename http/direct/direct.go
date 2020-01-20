package direct

import (
	"net/http"

	"github.com/wzshiming/pipe/http/template"
)

type Direct struct {
	code int
	body template.Format
}

func NewDirect(code int, body template.Format) *Direct {
	return &Direct{code, body}
}

func (d *Direct) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if d.code != 0 {
		rw.WriteHeader(d.code)
	}
	d.body.Format(rw, r)
}
