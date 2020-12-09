package template

import (
	"net/http"
)

var (
	NotFoundText    = "404 not found {{.Scheme}}://{{.Host}}{{.RequestURI}}"
	fmtNotFound, _  = NewFormat(NotFoundText)
	NotFoundHandler = http.HandlerFunc(NotFound)
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, fmtNotFound.FormatString(r), http.StatusNotFound)
}
