package forward

import (
	"context"

	"github.com/wzshiming/pipe/components/stream"
	"github.com/wzshiming/pipe/internal/logger"
	"github.com/wzshiming/pipe/internal/tunnel"
)

type Forward struct {
	dialer stream.Dialer
}

func NewForward(dialer stream.Dialer) *Forward {
	return &Forward{
		dialer: dialer,
	}
}

func (f *Forward) ServeStream(ctx context.Context, stm stream.Stream) {
	conn, err := f.dialer.DialStream(ctx)
	if err != nil {
		logger.Errorln(err)
		return
	}
	defer conn.Close()
	err = tunnel.Tunnel(ctx, stm, conn)
	if err != nil {
		logger.Errorln(err)
		return
	}
}
