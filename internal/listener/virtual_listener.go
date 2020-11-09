package listener

import (
	"context"
	"fmt"
	"io"
	"net"
	"sync"
	"sync/atomic"
)

type virtualListenerManager struct {
	port    map[int]*virtualListener
	address map[string]*virtualListener
	mut     sync.RWMutex
}

func newVirtualListenerManager() *virtualListenerManager {
	return &virtualListenerManager{
		port:    map[int]*virtualListener{},
		address: map[string]*virtualListener{},
	}
}

var ipv4zero = net.IPv4zero

func (v *virtualListenerManager) Listen(ctx context.Context, network, address string) (net.Listener, error) {
	addr, err := net.ResolveTCPAddr(network, address)
	if err != nil {
		return nil, err
	}

	listener := newVirtualListener(v, addr)

	var old io.Closer
	defer func() {
		if old != nil {
			old.Close()
		}
	}()
	v.mut.Lock()
	defer v.mut.Unlock()
	if addr.IP.Equal(ipv4zero) {
		l, ok := v.port[addr.Port]
		if ok {
			old = l
		}
		v.port[addr.Port] = listener
	} else {
		a := addr.String()
		l, ok := v.address[a]
		if ok {
			old = l
		}
		v.address[a] = listener
	}
	return listener, nil
}

func (v *virtualListenerManager) DialContext(ctx context.Context, network, address string) (net.Conn, error) {
	addr, err := net.ResolveTCPAddr(network, address)
	if err != nil {
		return nil, err
	}

	v.mut.RLock()
	defer v.mut.RUnlock()
	l, ok := v.address[addr.String()]
	if ok {
		return l.Conn(virtualAddr{})
	}
	l, ok = v.port[addr.Port]
	if ok {
		return l.Conn(virtualAddr{})
	}
	return nil, fmt.Errorf("couldn't connect to virtual server %s://%s", network, address)
}

func (v *virtualListenerManager) close(listener *virtualListener) error {
	address := listener.Addr()
	addr, err := net.ResolveTCPAddr(address.Network(), address.String())
	if err != nil {
		return err
	}

	v.mut.Lock()
	defer v.mut.Unlock()
	if addr.IP.Equal(ipv4zero) {
		l, ok := v.port[addr.Port]
		if ok && l == listener {
			delete(v.port, addr.Port)
		}
	} else {
		a := addr.String()
		l, ok := v.address[a]
		if ok && l == listener {
			delete(v.address, a)
		}
	}
	return nil
}

type virtualListener struct {
	parent     *virtualListenerManager
	serverAddr net.Addr
	ch         chan net.Conn
	isClose    uint32
}

func newVirtualListener(parent *virtualListenerManager, serverAddr net.Addr) *virtualListener {
	return &virtualListener{
		parent:     parent,
		serverAddr: serverAddr,
		ch:         make(chan net.Conn),
	}
}

func (l *virtualListener) Accept() (net.Conn, error) {
	conn, ok := <-l.ch
	if !ok {
		return nil, ErrNetClosing
	}
	return conn, nil
}

func (l *virtualListener) Close() error {
	if atomic.CompareAndSwapUint32(&l.isClose, 0, 1) {
		close(l.ch)
		if l.parent != nil {
			l.parent.close(l)
		}
	}
	return nil
}

func (l *virtualListener) Addr() net.Addr {
	return l.serverAddr
}

func (l *virtualListener) Conn(clientAddr net.Addr) (net.Conn, error) {
	if atomic.LoadUint32(&l.isClose) == 1 {
		return nil, ErrNetClosing
	}
	c, s := net.Pipe()
	s = &pipeConn{
		Conn:       s,
		remoteAddr: clientAddr,
		localAddr:  l.serverAddr,
	}
	c = &pipeConn{
		Conn:       c,
		remoteAddr: l.serverAddr,
		localAddr:  clientAddr,
	}
	l.ch <- s
	return c, nil
}

type pipeConn struct {
	net.Conn
	localAddr  net.Addr
	remoteAddr net.Addr
}

func (c *pipeConn) LocalAddr() net.Addr {
	return c.localAddr
}

func (c *pipeConn) RemoteAddr() net.Addr {
	return c.remoteAddr
}

type virtualAddr struct {
}

func (virtualAddr) Network() string {
	return "tcp"
}

func (virtualAddr) String() string {
	return "localhost:0"
}
