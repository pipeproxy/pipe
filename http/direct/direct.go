package direct

import (
	"net/http"
)

type Direct struct {
	code int
	body []byte
}

func NewDirect(code int, body []byte) *Direct {
	return &Direct{code, body}
}

func (d *Direct) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if d.code != 0 {
		rw.WriteHeader(d.code)
	}
	rw.Write(d.body)
}
