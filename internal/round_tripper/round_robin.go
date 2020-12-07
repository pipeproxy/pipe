package round_tripper

import (
	"net/http"
	"sync/atomic"
)

type RoundRobin struct {
	handlers []http.RoundTripper
	count    uint64
}

func NewRoundRobin(handlers []http.RoundTripper) *Random {
	return &Random{handlers: handlers}
}

func (r *RoundRobin) RoundTrip(req *http.Request) (*http.Response, error) {
	return r.handlers[int(atomic.AddUint64(&r.count, 1)-1)%len(r.handlers)].RoundTrip(req)
}
