package poller

import (
	"math/rand"
	"net/http"
)

type Random struct {
	handlers []http.Handler
}

func NewRandom(handlers []http.Handler) *Random {
	return &Random{handlers: handlers}
}

func (r *Random) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	r.handlers[rand.Int63n(int64(len(r.handlers)))].ServeHTTP(rw, req)
}
