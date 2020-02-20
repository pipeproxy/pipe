package dialer

import (
	"context"

	"github.com/wzshiming/pipe/configure/alias"
	"github.com/wzshiming/pipe/pipe/common/load"
	"github.com/wzshiming/pipe/pipe/stream"
)

func init() {
	var dialer Dialer
	alias.Register("stream.Dialer", &dialer)
	load.Register(&dialer)
}

type Dialer interface {
	DialStream(ctx context.Context) (stream.Stream, error)
}
