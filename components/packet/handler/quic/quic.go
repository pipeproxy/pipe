package quic

import (
	"context"
	"net"
	"sync"

	quic "github.com/lucas-clemente/quic-go"
	"github.com/pipeproxy/pipe/internal/listener"
	"github.com/pipeproxy/pipe/internal/logger"
)

type Listener struct {
	quic.Listener
	ch     chan net.Conn
	ctx    context.Context
	cancel func()
	once   sync.Once
}

func NewListener(ctx context.Context, listener quic.Listener) *Listener {
	l := &Listener{
		Listener: listener,
		ctx:      ctx,
	}
	return l
}

func (l *Listener) init() {
	l.ctx, l.cancel = context.WithCancel(l.ctx)
	l.ch = make(chan net.Conn)
	go l.accept()
}

func (l *Listener) accept() {
	for {
		sess, err := l.Listener.Accept(l.ctx)
		if err != nil {
			logger.Errorf("accept session failed: %v", err)
			return
		}

		go l.acceptStream(sess)
	}
}

func (l *Listener) acceptStream(sess quic.Session) {
	for {
		stm, err := sess.AcceptStream(l.ctx)
		if err != nil {
			logger.Errorf("accept stream failed: %v", err)
			return
		}

		conn := &streamWarp{streamAddr: sess, Stream: stm}
		l.ch <- conn
	}
}

func (l *Listener) Accept() (net.Conn, error) {
	l.once.Do(l.init)
	conn, ok := <-l.ch
	if !ok || conn == nil {
		return nil, listener.ErrNetClosing
	}
	return conn, nil
}

func (l *Listener) Close() error {
	l.cancel()
	close(l.ch)
	return nil
}

type streamWarp struct {
	streamAddr
	quic.Stream
}

type streamAddr interface {
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
}
