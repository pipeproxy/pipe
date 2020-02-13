package forward

import (
	"context"

	"github.com/wzshiming/pipe/dialer"
	"github.com/wzshiming/pipe/internal/joinio"
	"github.com/wzshiming/pipe/stream"
)

type Forward struct {
	dialer dialer.Dialer
}

func NewForward(dialer dialer.Dialer) *Forward {
	return &Forward{
		dialer: dialer,
	}
}

func (f *Forward) ServeStream(ctx context.Context, stm stream.Stream) {
	conn, err := f.dialer.Dial(ctx)
	if err != nil {
		return
	}
	defer conn.Close()
	joinio.BothCopy(stm, conn)
}
