package http

import (
	"io"
	"net"
)

type Listener struct {
	ch chan net.Conn
}

func NewListener() *Listener {
	return &Listener{
		ch: make(chan net.Conn),
	}
}

// Send a conn to listener.
func (l *Listener) Send(conn net.Conn) error {
	l.ch <- conn
	return nil
}

// Accept waits for and returns the next connection to the listener.
func (l *Listener) Accept() (net.Conn, error) {
	conn, ok := <-l.ch
	if !ok {
		return nil, io.ErrClosedPipe
	}
	return conn, nil
}

// Close closes the listener.
// Any blocked Accept operations will be unblocked and return errors.
func (l *Listener) Close() error {
	close(l.ch)
	return nil
}

// Addr returns the listener's network address.
func (l *Listener) Addr() net.Addr {
	return noneAddr{}
}

type noneAddr struct {
}

func (noneAddr) Network() string {
	return "none"
}
func (noneAddr) String() string {
	return "none"
}
