package stream

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
	"sync/atomic"

	"github.com/wzshiming/pipe/pipe/stream"
	"github.com/wzshiming/pipe/pipe/stream/listener"
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
			v.Shutdown()
			dk = append(dk, k)
		}
	}

	for _, k := range dk {
		log.Printf("[INFO] Close listen to %s", k)
		delete(cache, k)
	}
}

func Listen(ctx context.Context, network, address string) (listener.StreamListener, error) {
	mut.Lock()
	defer mut.Unlock()
	key := fmt.Sprintf("%s://%s", network, address)
	n, ok := cache[key]
	if ok {
		log.Printf("[INFO] Relisten to %s", key)
		return n.Listener(ctx), nil
	}

	log.Printf("[INFO] ListenStream to %s", key)
	l, err := net.Listen(network, address)
	if err != nil {
		return nil, err
	}
	n = newHub(l)
	cache[key] = n
	return n.Listener(ctx), nil
}

type Hub struct {
	listener listener.StreamListener
	ch       chan stream.Stream
	exit     chan struct{}
	size     int32
}

func newHub(listener listener.StreamListener) *Hub {
	m := &Hub{
		listener: listener,
		ch:       make(chan stream.Stream),
		exit:     make(chan struct{}),
	}
	go m.run()
	return m
}

func (h *Hub) run() {
	for {
		conn, err := h.listener.Accept()
		if err != nil {
			log.Printf("[ERROR] accept error %s", err)
			return
		}
		select {
		case h.ch <- conn:
		case <-h.exit:
			conn.Close()
		}
	}
}

func (h *Hub) Shutdown() error {
	err := h.listener.Close()
	close(h.exit)
	close(h.ch)
	return err
}

func (h *Hub) Listener(ctx context.Context) *Listener {
	l := &Listener{
		ctx:  ctx,
		hub:  h,
		exit: make(chan struct{}),
	}
	atomic.AddInt32(&h.size, 1)
	return l
}

type Listener struct {
	ctx       context.Context
	closeOnce sync.Once
	exit      chan struct{}
	hub       *Hub
}

func (l *Listener) Accept() (stream.Stream, error) {
	select {
	case <-l.ctx.Done():
		l.Close()
		err := l.ctx.Err()
		if err != nil {
			return nil, err
		}
		return nil, io.ErrClosedPipe
	case conn := <-l.hub.ch:
		if conn == nil {
			l.Close()
			return nil, io.ErrClosedPipe
		}
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

// Addr returns the listener's stream address.
func (l *Listener) Addr() net.Addr {
	return l.hub.listener.Addr()
}
