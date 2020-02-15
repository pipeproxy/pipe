package dialer

import (
	"context"

	"github.com/wzshiming/pipe/pipe/stream"
)

type Dialer interface {
	DialStream(ctx context.Context) (stream.Stream, error)
}
