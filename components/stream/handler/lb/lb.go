package lb

import (
	"context"

	"github.com/pipeproxy/pipe/components/balance"
	"github.com/pipeproxy/pipe/components/stream"
)

type LB struct {
	policy   balance.Policy
	handlers []stream.Handler
}

func NewLB(policy balance.Policy, handlers []stream.Handler) *LB {
	return &LB{policy: policy, handlers: handlers}
}

func (l *LB) ServeStream(ctx context.Context, stm stream.Stream) {
	l.policy.InUse(uint64(len(l.handlers)), func(i uint64) {
		l.handlers[i].ServeStream(ctx, stm)
	})
}
