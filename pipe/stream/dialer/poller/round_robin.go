package poller

import (
	"context"
	"sync/atomic"

	"github.com/wzshiming/pipe/pipe/stream"
	"github.com/wzshiming/pipe/pipe/stream/dialer"
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
