package round_tripper

import (
	"net/http"

	"github.com/wzshiming/pipe/components/common/types"
)

func init() {
	var roundTripper RoundTripper
	types.Register(&roundTripper)
}

type RoundTripper = http.RoundTripper
