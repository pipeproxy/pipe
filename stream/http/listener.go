package http

import (
	"io"
	"net"

	"github.com/wzshiming/pipe/stream"
)

type singleConnListener struct {
	conn stream.Stream
}

func (l *singleConnListener) Accept() (net.Conn, error) {
	conn := l.conn
	if conn == nil {
		return nil, io.ErrClosedPipe
	}
	l.conn = nil
	return conn, nil
}

func (*singleConnListener) Close() error {
	return nil
}

func (*singleConnListener) Addr() net.Addr {
	return nil
}
