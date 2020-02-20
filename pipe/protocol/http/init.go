package http

import (
	"net/http"

	"github.com/wzshiming/pipe/configure/alias"
	"github.com/wzshiming/pipe/pipe/common/load"
)

func init() {
	var handler Handler
	alias.Register("http.Handler", &handler)
	load.Register(&handler)
}

type Handler = http.Handler
