package round_tripper

import (
	"net/http"

	"github.com/wzshiming/pipe/configure/alias"
)

func init() {
	var roundTripper RoundTripper
	alias.Register("http.RoundTripper", &roundTripper)
}

type RoundTripper = http.RoundTripper
