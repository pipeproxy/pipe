package network

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/wzshiming/pipe/components/stream"
	"github.com/wzshiming/pipe/components/stream/listener"
	"github.com/wzshiming/pipe/internal/logger"
)

var (
	mut   = sync.Mutex{}
	cache = map[string]*Hub{}
)

func closeExcess() {
	mut.Lock()
	defer mut.Unlock()
	dk := []string{}
	for k, v := range cache {
		if atomic.LoadInt32(&v.size) == 0 {
			v.shutdown()
			dk = append(dk, k)
		}
	}

	for _, k := range dk {
		logger.Infof("Close listen to %s", k)
		delete(cache, k)
	}
}

var (
	chClose     chan struct{}
	chCloseOnce sync.Once
)

func channelClose() {
	chCloseOnce.Do(func() {
		chClose = make(chan struct{}, 1)
		go func() {
		loop:
			for range chClose {
				for {
					select {
					case <-chClose:
					case <-time.After(time.Second / 10):
						closeExcess()
						continue loop
					}
				}
			}
		}()
	})
	select {
	case chClose <- struct{}{}:
	default:
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
	size     int32
}

func newHub(listener listener.StreamListener) *Hub {
	m := &Hub{
		listener: listener,
		ch:       make(chan stream.Stream),
	}
	go m.run()
	return m
}

func (h *Hub) run() {
	for atomic.LoadInt32(&h.size) != -1 {
		conn, err := h.listener.Accept()
		if err != nil {
			if !IsClosedConnError(err) {
				logger.Errorf("accept error %#v", err)
			}
			break
		}
		h.ch <- conn
	}
	close(h.ch)
}

func (h *Hub) shutdown() error {
	if !atomic.CompareAndSwapInt32(&h.size, 0, -1) {
		return nil
	}

	return h.listener.Close()
}

func (h *Hub) Listener(ctx context.Context) *Listener {
	var once sync.Once
	ctx, cancel := context.WithCancel(ctx)
	l := &Listener{
		ctx: ctx,
		hub: h,
		close: func() {
			once.Do(func() {
				cancel()
				if atomic.AddInt32(&h.size, -1) == 0 {
					channelClose()
				}
			})
		},
	}
	atomic.AddInt32(&h.size, 1)
	return l
}

type Listener struct {
	ctx   context.Context
	close func()
	hub   *Hub
}

func (l *Listener) Accept() (stream.Stream, error) {
	select {
	case <-l.ctx.Done():
		l.Close()
		return nil, ErrNetClosing
	case conn, ok := <-l.hub.ch:
		if !ok || conn == nil {
			l.Close()
			return nil, ErrNetClosing
		}
		return conn, nil
	}
}

// Close closes the listener.
func (l *Listener) Close() error {
	l.close()
	return nil
}

// Addr returns the listener's stream address.
func (l *Listener) Addr() net.Addr {
	return l.hub.listener.Addr()
}

// Code adapted from net/http

// ErrNetClosing is returned when a network descriptor is used after
// it has been closed. Keep this string consistent because of issue
// #4373: since historically programs have not been able to detect
// this error, they look for the string.
var ErrNetClosing = errors.New("use of closed network connection")

// IsClosedConnError reports whether err is an error from use of a closed
// network connection.
func IsClosedConnError(err error) bool {
	if err == nil {
		return false
	}

	// TODO: remove this string search and be more like the Windows
	// case below. That might involve modifying the standard library
	// to return better error types.
	if err == ErrNetClosing || strings.Contains(err.Error(), ErrNetClosing.Error()) {
		return true
	}

	// TODO(bradfitz): x/tools/cmd/bundle doesn't really support
	// build tags, so I can't make an http2_windows.go file with
	// Windows-specific stuff. Fix that and move this, once we
	// have a way to bundle this into std's net/http somehow.
	if runtime.GOOS == "windows" {
		if oe, ok := err.(*net.OpError); ok && oe.Op == "read" {
			if se, ok := oe.Err.(*os.SyscallError); ok && se.Syscall == "wsarecv" {
				const WSAECONNABORTED = 10053
				const WSAECONNRESET = 10054
				if n := errno(se.Err); n == WSAECONNRESET || n == WSAECONNABORTED {
					return true
				}
			}
		}
	}
	return false
}

// errno returns v's underlying uintptr, else 0.
//
// TODO: remove this helper function once http2 can use build
// tags. See comment in isClosedConnError.
func errno(v error) uintptr {
	if rv := reflect.ValueOf(v); rv.Kind() == reflect.Uintptr {
		return uintptr(rv.Uint())
	}
	return 0
}
