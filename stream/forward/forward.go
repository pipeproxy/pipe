package forward

import (
	"context"
	"io"

	"github.com/wzshiming/pipe/dialer"
	"github.com/wzshiming/pipe/internal/pool"
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

	buf1 := pool.Buffer.Get()
	buf2 := pool.Buffer.Get()
	go io.CopyBuffer(stm, conn, buf1[:])
	io.CopyBuffer(conn, stm, buf2[:])

	pool.Buffer.Put(buf1)
	pool.Buffer.Put(buf2)

}
