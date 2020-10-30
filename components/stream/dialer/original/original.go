package original

import (
	"context"
	"errors"
	"net"

	"github.com/mikioh/tcp"
	svc_stream "github.com/wzshiming/pipe/components/service/stream"
	"github.com/wzshiming/pipe/components/stream"
)

type Original struct {
	dialer net.Dialer
}

func NewOriginal() stream.Dialer {
	return &Original{}
}

func (d *Original) DialStream(ctx context.Context) (stream.Stream, error) {
	s, ok := svc_stream.GetRawStreamWithContext(ctx)
	if !ok {
		return nil, errors.New("unable to get raw stream")
	}
	c, err := tcp.NewConn(s)
	if err != nil {
		return nil, err
	}
	addr, err := c.OriginalDst()
	if err != nil {
		return nil, err
	}
	return d.dialer.DialContext(ctx, addr.Network(), addr.String())
}
