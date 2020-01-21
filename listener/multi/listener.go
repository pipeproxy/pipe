package multi

import (
	"io"
	"log"
	"net"
)

type multiListener struct {
	ch    chan net.Conn
	multi []net.Listener
}

func newMultiListener(multi []net.Listener) (*multiListener, error) {
	m := &multiListener{
		ch:    make(chan net.Conn),
		multi: multi,
	}
	err := m.init()
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (l *multiListener) init() error {
	for _, listener := range l.multi {
		go l.start(listener)
	}
	return nil
}

func (l *multiListener) start(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("[ERROR] start accept: %s", err)
			return
		}
		l.ch <- conn
	}
}

// Accept waits for and returns the next connection to the listener.
func (l *multiListener) Accept() (net.Conn, error) {
	conn, ok := <-l.ch
	if !ok {
		return nil, io.ErrClosedPipe
	}
	return conn, nil
}

// Close closes the listener.
// Any blocked Accept operations will be unblocked and return errors.
func (l *multiListener) Close() error {
	for _, listener := range l.multi {
		err := listener.Close()
		if err != nil {
			log.Printf("[ERROR] close listen: %s", err)
		}
	}
	close(l.ch)
	return nil
}

// Addr returns the listener's network address.
func (l *multiListener) Addr() net.Addr {
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
