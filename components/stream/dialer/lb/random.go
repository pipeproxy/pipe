package lb

import (
	"context"
	"math/rand"

	"github.com/wzshiming/pipe/components/stream"
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
