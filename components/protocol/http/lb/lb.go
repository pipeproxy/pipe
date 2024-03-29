package lb

import (
	"net/http"

	"github.com/pipeproxy/pipe/components/balance"
)

type LB struct {
	policy   balance.Policy
	handlers []http.Handler
}

func NewLB(policy balance.Policy, handlers []http.Handler) *LB {
	policy.Init(uint64(len(handlers)))
	return &LB{policy: policy, handlers: handlers}
}

func (l *LB) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	l.policy.InUse(func(i uint64) {
		l.handlers[i].ServeHTTP(rw, r)
	})
}
