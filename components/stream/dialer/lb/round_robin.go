package lb

import (
	"context"
	"net/http"
	"sync/atomic"

	"github.com/pipeproxy/pipe/components/stream"
	"github.com/pipeproxy/pipe/internal/round_tripper"
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

func (r *RoundRobin) RoundTripper() http.RoundTripper {
	return round_tripper.NewRoundRobin(round_tripper.RoundTripperList(r.handlers))
}
