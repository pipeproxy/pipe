package poller

import (
	"context"
	"sync/atomic"

	"github.com/wzshiming/pipe/dialer"

	"github.com/wzshiming/pipe/stream"
)

type RoundRobin struct {
	dialers []dialer.Dialer
	count   uint64
}

func NewRoundRobin(dialers []dialer.Dialer) *Random {
	return &Random{dialers: dialers}
}

func (r *RoundRobin) Dial(ctx context.Context) (stream.Stream, error) {
	return r.dialers[int(atomic.AddUint64(&r.count, 1))%len(r.dialers)].Dial(ctx)
}
