package network

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
	"sync/atomic"
)

var (
	mut   = sync.Mutex{}
	cache = map[string]*Hub{}
)

func CloseExcess() {
	mut.Lock()
	defer mut.Unlock()
	dk := []string{}
	for k, v := range cache {
		if atomic.LoadInt32(&v.size) == 0 {
			v.Close()
			dk = append(dk, k)
		}
	}

	for _, k := range dk {
		log.Printf("[INFO] Close listen to %s", k)
		delete(cache, k)
	}
}

func Listen(ctx context.Context, network, address string) (net.Listener, error) {
	mut.Lock()
	defer mut.Unlock()
	key := fmt.Sprintf("%s://%s", network, address)
	n, ok := cache[key]
	if ok {
		log.Printf("[INFO] Relisten to %s", key)
		return n.Listener(ctx), nil
	}

	log.Printf("[INFO] Listen to %s", key)
	l, err := net.Listen(network, address)
	if err != nil {
		return nil, err
	}
	n = newListener(l)
	cache[key] = n
	return n.Listener(ctx), nil
}

type Hub struct {
	listener net.Listener
	ch       chan net.Conn
	size     int32
}

func newListener(listener net.Listener) *Hub {
	m := &Hub{
		listener: listener,
		ch:       make(chan net.Conn),
	}
	go m.run()
	return m
}

func (h *Hub) run() {
	for {
		conn, err := h.listener.Accept()
		if err != nil {
			return
		}
		h.ch <- conn
	}
}

func (h *Hub) Close() error {
	h.listener.Close()
	close(h.ch)

	return nil
}

func (h *Hub) Listener(ctx context.Context) *Listener {
	l := &Listener{
		ctx:  ctx,
		hub:  h,
		ch:   h.ch,
		exit: make(chan struct{}),
	}
	atomic.AddInt32(&h.size, 1)
	return l
}

type Listener struct {
	ctx       context.Context
	closeOnce sync.Once
	ch        chan net.Conn
	exit      chan struct{}
	hub       *Hub
}

func (l *Listener) Accept() (net.Conn, error) {
	select {
	case <-l.ctx.Done():
		l.Close()
		return nil, l.ctx.Err()
	case conn := <-l.ch:
		return conn, nil
	case <-l.exit:
		return nil, io.ErrClosedPipe
	}
}

// Close closes the listener.
func (l *Listener) Close() error {
	l.closeOnce.Do(func() {
		atomic.AddInt32(&l.hub.size, -1)
		close(l.exit)
	})
	return nil
}

// Addr returns the listener's network address.
func (l *Listener) Addr() net.Addr {
	return l.hub.listener.Addr()
}
