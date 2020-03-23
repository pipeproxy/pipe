package http

import (
	"net/http"

	"github.com/wzshiming/pipe/components/common/types"
)

func init() {
	var handler Handler
	types.Register(&handler)
}

type Handler = http.Handler
