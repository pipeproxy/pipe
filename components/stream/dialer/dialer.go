package dialer

import (
	"context"

	"github.com/wzshiming/pipe/components/common/types"
	"github.com/wzshiming/pipe/components/stream"
)

func init() {
	var dialer Dialer
	types.Register(&dialer)
}

type Dialer interface {
	DialStream(ctx context.Context) (stream.Stream, error)
}
