package dialer

import (
	"context"

	"github.com/wzshiming/pipe/configure/alias"
	"github.com/wzshiming/pipe/pipe/stream"
)

func init() {
	var dialer Dialer
	alias.Register("stream.Dialer", &dialer)
}

type Dialer interface {
	DialStream(ctx context.Context) (stream.Stream, error)
}
