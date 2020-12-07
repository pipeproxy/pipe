package round_tripper

import (
	"math/rand"
	"net/http"
)

type Random struct {
	handlers []http.RoundTripper
}

func NewRandom(handlers []http.RoundTripper) *Random {
	return &Random{handlers: handlers}
}

func (r *Random) RoundTrip(req *http.Request) (*http.Response, error) {
	return r.handlers[rand.Int63n(int64(len(r.handlers)))].RoundTrip(req)
}
