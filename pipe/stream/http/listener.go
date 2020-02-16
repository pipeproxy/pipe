package http

import (
	"net"
	"net/http"

	"github.com/wzshiming/pipe/pipe/stream"
)

type singleConnListener struct {
	addr net.Addr
	ch   chan stream.Stream
}

func newSingleConnListener(conn stream.Stream) *singleConnListener {
	ch := make(chan stream.Stream, 1)
	ch <- conn
	return &singleConnListener{
		addr: conn.LocalAddr(),
		ch:   ch,
	}
}

func (l *singleConnListener) Accept() (stream.Stream, error) {
	conn := <-l.ch
	if conn == nil {
		return nil, http.ErrServerClosed
	}
	return &connCloser{
		l:    l,
		Conn: conn,
	}, nil
}

func (l *singleConnListener) shutdown() error {
	close(l.ch)
	return nil
}

func (l *singleConnListener) Close() error {
	return nil
}

func (l *singleConnListener) Addr() net.Addr {
	return l.addr
}

type connCloser struct {
	l *singleConnListener
	stream.Stream
}

func (c *connCloser) Close() error {
	return c.l.shutdown()
}
