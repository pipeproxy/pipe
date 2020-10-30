package lb

import (
	"context"
	"sync/atomic"

	"github.com/pipeproxy/pipe/components/stream"
)

type RoundRobin struct {
	handlers []stream.Handler
	count    uint64
}

func NewRoundRobin(handlers []stream.Handler) *Random {
	return &Random{handlers: handlers}
}

func (r *RoundRobin) ServeStream(ctx context.Context, stm stream.Stream) {
	r.handlers[int(atomic.AddUint64(&r.count, 1))%len(r.handlers)].ServeStream(ctx, stm)
}
