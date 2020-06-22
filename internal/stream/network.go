package stream

import (
	"context"
	"fmt"
	"io"
	"net"
	"sync"
	"sync/atomic"

	"github.com/wzshiming/pipe/components/stream"
	"github.com/wzshiming/pipe/components/stream/listener"
	"github.com/wzshiming/pipe/internal/logger"
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
		logger.Infof("Close listen to %s", k)
		delete(cache, k)
	}
}

func ListenList() []string {
	mut.Lock()
	defer mut.Unlock()
	dk := []string{}
	for k := range cache {
		dk = append(dk, k)
	}
	return dk
}

func Listen(ctx context.Context, network, address string) (listener.StreamListener, error) {
	mut.Lock()
	defer mut.Unlock()

	if _, port, _ := net.SplitHostPort(address); port != "" && port != "0" {
		key := fmt.Sprintf("%s://%s", network, address)
		n, ok := cache[key]
		if ok {
			logger.Infof("Relisten to %s", key)
			return n.Listener(ctx), nil
		}
	}

	var lc net.ListenConfig
	l, err := lc.Listen(ctx, network, address)
	if err != nil {
		return nil, err
	}

	address = sameAddress(address, l.Addr().String())
	key := fmt.Sprintf("%s://%s", network, address)
	logger.Infof("Listen to %s", key)
	n := newHub(l)
	cache[key] = n
	return n.Listener(ctx), nil
}

func sameAddress(a1, a2 string) string {
	host1, port1, err := net.SplitHostPort(a1)
	if err != nil {
		return a1
	}

	switch host1 {
	case "0.0.0.0", "[::]":
		host1 = ""
	}

	switch port1 {
	case "0":
		_, port2, err := net.SplitHostPort(a2)
		if err != nil {
			return a1
		}
		port1 = port2
	}
	return fmt.Sprintf("%s:%s", host1, port1)
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
			logger.Errorf("accept error %s", err)
			return
		}
		if atomic.LoadInt32(&h.size) == 0 {
			conn.Close()
			return
		}
		select {
		case h.ch <- conn:
		case <-h.exit:
			conn.Close()
			return
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
		return nil, io.ErrClosedPipe
	case conn, ok := <-l.hub.ch:
		if !ok || conn == nil {
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
