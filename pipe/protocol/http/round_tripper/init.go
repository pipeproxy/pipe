package round_tripper

import (
	"net/http"

	"github.com/wzshiming/pipe/configure/alias"
	"github.com/wzshiming/pipe/pipe/common/load"
)

func init() {
	var roundTripper RoundTripper
	alias.Register("http.RoundTripper", &roundTripper)
	load.Register(&roundTripper)
}

type RoundTripper = http.RoundTripper
