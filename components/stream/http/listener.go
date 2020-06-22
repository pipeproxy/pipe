package http

import (
	"net"
	"net/http"
	"sync"

	"github.com/wzshiming/pipe/components/stream"
)

type singleConnListener struct {
	addr net.Addr
	ch   chan stream.Stream
	once sync.Once
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
	conn, ok := <-l.ch
	if !ok || conn == nil {
		return nil, http.ErrServerClosed
	}
	return &connCloser{
		l:      l,
		Stream: conn,
	}, nil
}

func (l *singleConnListener) shutdown() error {
	l.once.Do(func() {
		close(l.ch)
	})
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
	c.l.shutdown()
	return c.Stream.Close()
}
