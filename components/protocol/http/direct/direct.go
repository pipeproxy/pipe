package direct

import (
	"net/http"

	"github.com/pipeproxy/pipe/internal/http/template"
	"github.com/wzshiming/logger"
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
	} else {
		rw.WriteHeader(http.StatusOK)
	}
	if d.body != nil {
		err := d.body.Format(rw, r)
		if err != nil {
			logger.FromContext(r.Context()).Error(err, "Format")
		}
	}
}
