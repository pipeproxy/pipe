package poller

import (
	"context"
	"math/rand"

	"github.com/wzshiming/pipe/pipe/stream"
	"github.com/wzshiming/pipe/pipe/stream/dialer"
)

type Random struct {
	dialers []dialer.Dialer
}

func NewRandom(dialers []dialer.Dialer) *Random {
	return &Random{dialers: dialers}
}

func (r *Random) DialStream(ctx context.Context) (stream.Stream, error) {
	return r.dialers[rand.Int63n(int64(len(r.dialers)))].DialStream(ctx)
}
