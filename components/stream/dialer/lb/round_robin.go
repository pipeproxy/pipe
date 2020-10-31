package lb

import (
	"context"
	"sync/atomic"

	"github.com/pipeproxy/pipe/components/stream"
)

type RoundRobin struct {
	handlers []stream.Dialer
	count    uint64
}

func NewRoundRobin(handlers []stream.Dialer) *Random {
	return &Random{handlers: handlers}
}

func (r *RoundRobin) DialStream(ctx context.Context) (stream.Stream, error) {
	return r.handlers[int(atomic.AddUint64(&r.count, 1)-1)%len(r.handlers)].DialStream(ctx)
}
