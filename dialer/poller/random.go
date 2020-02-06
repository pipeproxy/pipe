package poller

import (
	"context"
	"math/rand"

	"github.com/wzshiming/pipe/dialer"
	"github.com/wzshiming/pipe/stream"
)

type Random struct {
	dialers []dialer.Dialer
}

func NewRandom(dialers []dialer.Dialer) *Random {
	return &Random{dialers: dialers}
}

func (r *Random) Dial(ctx context.Context) (stream.Stream, error) {
	return r.dialers[rand.Int63n(int64(len(r.dialers)))].Dial(ctx)
}
