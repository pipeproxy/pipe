package redirect

import (
	"net/http"

	"github.com/wzshiming/pipe/components/protocol/http/template"
)

type Redirect struct {
	code     int
	location template.Format
}

func NewRedirect(code int, location template.Format) *Redirect {
	return &Redirect{
		code:     code,
		location: location,
	}
}

func (h *Redirect) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	location := h.location.FormatString(r)
	http.Redirect(rw, r, location, h.code)
}
