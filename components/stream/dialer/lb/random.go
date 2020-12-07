package lb

import (
	"context"
	"math/rand"
	"net/http"

	"github.com/pipeproxy/pipe/components/stream"
	"github.com/pipeproxy/pipe/internal/round_tripper"
)

type Random struct {
	handlers []stream.Dialer
}

func NewRandom(handlers []stream.Dialer) *Random {
	return &Random{handlers: handlers}
}

func (r *Random) DialStream(ctx context.Context) (stream.Stream, error) {
	return r.handlers[rand.Int63n(int64(len(r.handlers)))].DialStream(ctx)
}

func (r *Random) RoundTripper() http.RoundTripper {
	return round_tripper.NewRandom(round_tripper.RoundTripperList(r.handlers))
}
