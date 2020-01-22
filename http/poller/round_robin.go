package poller

import (
	"net/http"
	"sync/atomic"
)

type RoundRobin struct {
	handlers []http.Handler
	count    uint64
}

func NewRoundRobin(handlers []http.Handler) *Random {
	return &Random{handlers: handlers}
}

func (r *RoundRobin) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	r.handlers[int(atomic.AddUint64(&r.count, 1))%len(r.handlers)].ServeHTTP(rw, req)
}
