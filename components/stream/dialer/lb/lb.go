package lb

import (
	"context"
	"strings"

	"github.com/pipeproxy/pipe/components/balance"
	"github.com/pipeproxy/pipe/components/stream"
)

type LB struct {
	policy  balance.Policy
	dialers []stream.Dialer
	name    string
}

func NewLB(policy balance.Policy, dialers []stream.Dialer) *LB {
	policy.Init(uint64(len(dialers)))
	l := &LB{policy: policy, dialers: dialers}
	l.name = l.getName()
	return l
}

func (l *LB) DialStream(ctx context.Context) (stm stream.Stream, err error) {
	l.policy.InUse(func(i uint64) {
		stm, err = l.dialers[i].DialStream(ctx)
	})
	return
}

func (l *LB) Targets() []stream.Dialer {
	return l.dialers
}

func (l *LB) Policy() balance.Policy {
	return l.policy.Clone()
}

func (l *LB) String() string {
	return l.name
}

func (l *LB) getName() string {
	strs := make([]string, 0, len(l.dialers))
	for _, dialer := range l.dialers {
		strs = append(strs, dialer.String())
	}
	return strings.Join(strs, ";")
}
