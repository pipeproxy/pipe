package http

import (
	"net/http"

	"github.com/wzshiming/pipe/configure/alias"
)

func init() {
	var handler Handler
	alias.Register("http.Handler", &handler)
}

type Handler = http.Handler
