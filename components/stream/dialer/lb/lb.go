package lb

import (
	"context"

	"github.com/pipeproxy/pipe/components/balance"
	"github.com/pipeproxy/pipe/components/stream"
)

type LB struct {
	policy  balance.Policy
	dialers []stream.Dialer
}

func NewLB(policy balance.Policy, dialers []stream.Dialer) *LB {
	return &LB{policy: policy, dialers: dialers}
}

func (l *LB) DialStream(ctx context.Context) (stm stream.Stream, err error) {
	l.policy.InUse(uint64(len(l.dialers)), func(i uint64) {
		stm, err = l.dialers[i].DialStream(ctx)
	})
	return
}

func (l *LB) Targets() (balance.PolicyEnum, []stream.Dialer) {
	return l.policy.Policy(), l.dialers
}
