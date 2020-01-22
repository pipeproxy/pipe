package http

import (
	"io"
	"net"
)

type singleConnListener struct {
	conn net.Conn
}

func (l *singleConnListener) Accept() (net.Conn, error) {
	conn := l.conn
	if conn == nil {
		return nil, io.ErrClosedPipe
	}
	l.conn = nil
	return conn, nil
}

func (l *singleConnListener) Close() error {
	return nil
}

func (l *singleConnListener) Addr() net.Addr {
	return nil
}
