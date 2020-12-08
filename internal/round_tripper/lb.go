package round_tripper

import (
	"net/http"

	"github.com/pipeproxy/pipe/components/balance"
)

type LB struct {
	policy   balance.Policy
	trippers []http.RoundTripper
}

func NewLB(policy balance.Policy, trippers []http.RoundTripper) *LB {
	policy.Init(uint64(len(trippers)))
	return &LB{policy: policy, trippers: trippers}
}

func (l *LB) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	l.policy.InUse(func(i uint64) {
		resp, err = l.trippers[i].RoundTrip(req)
	})
	return
}
