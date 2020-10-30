package lb

import (
	"context"
	"math/rand"

	"github.com/pipeproxy/pipe/components/stream"
)

type Random struct {
	handlers []stream.Handler
}

func NewRandom(handlers []stream.Handler) *Random {
	return &Random{handlers: handlers}
}

func (r *Random) ServeStream(ctx context.Context, stm stream.Stream) {
	r.handlers[rand.Int63n(int64(len(r.handlers)))].ServeStream(ctx, stm)
}
