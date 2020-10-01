package listener

import (
	"context"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/wzshiming/pipe/internal/logger"
)

// Listen returns a listener suitable.
func Listen(ctx context.Context, network, address string) (net.Listener, error) {
	err := ctx.Err()
	if err != nil {
		return nil, err
	}
	ln, err := listen(ctx, network, address)
	if err != nil {
		return nil, err
	}
	swapClose(ctx, ln)
	return ln, nil
}

func listen(ctx context.Context, network, address string) (*fakeCloseListener, error) {
	listenerMut.Lock()
	defer listenerMut.Unlock()

	key := buildKey(network, address)
	if _, port, _ := net.SplitHostPort(address); port != "" && port != "0" {
		if global, ok := listeners[key]; ok {
			atomic.AddInt32(&global.usage, 1)
			logger.Infof("Relisten to %s", key)
			return &fakeCloseListener{
				usage:       &global.usage,
				key:         key,
				Listener:    global.listener,
				deadline:    &global.deadline,
				deadlineMut: &global.deadlineMut,
			}, nil
		}
	}

	var listenConfig net.ListenConfig
	ln, err := listenConfig.Listen(ctx, network, address)
	if err != nil {
		return nil, err
	}
	address = sameAddress(address, ln.Addr().String())
	global := &globalListener{usage: 1, listener: ln}
	listeners[key] = global
	realKey := buildKey(network, address)
	if key != realKey {
		logger.Infof("Listen to %s (%s)", key, realKey)
	} else {
		logger.Infof("Listen to %s", key)
	}
	return &fakeCloseListener{
		usage:       &global.usage,
		key:         key,
		Listener:    ln,
		deadline:    &global.deadline,
		deadlineMut: &global.deadlineMut,
	}, nil
}

type fakeCloseListener struct {
	closed      int32
	usage       *int32
	key         string
	deadline    *bool       // protected by deadlineMut; global
	deadlineMut *sync.Mutex // global
	net.Listener
}

// Accept accepts connections until CloseSwap() is called.
func (f *fakeCloseListener) Accept() (net.Conn, error) {
	if atomic.LoadInt32(&f.closed) == 1 {
		return nil, ErrNetClosing
	}

	conn, err := f.Listener.Accept()
	if err == nil {
		return conn, nil
	}

	f.deadlineMut.Lock()
	if *f.deadline {
		switch ln := f.Listener.(type) {
		case *net.TCPListener:
			ln.SetDeadline(time.Time{})
		case *net.UnixListener:
			ln.SetDeadline(time.Time{})
		}
		*f.deadline = false
	}
	f.deadlineMut.Unlock()

	if IsClosedConnError(err) || IsAcceptTimeoutError(err) {
		return nil, ErrNetClosing
	}

	return nil, err
}

func (f *fakeCloseListener) Key() string {
	return f.key
}

func (f *fakeCloseListener) Close() error {
	if !atomic.CompareAndSwapInt32(&f.closed, 0, 1) {
		return nil
	}

	f.deadlineMut.Lock()
	if !*f.deadline {
		switch ln := f.Listener.(type) {
		case *net.TCPListener:
			ln.SetDeadline(time.Now().Add(-1 * time.Minute))
		case *net.UnixListener:
			ln.SetDeadline(time.Now().Add(-1 * time.Minute))
		}
		*f.deadline = true
	}
	f.deadlineMut.Unlock()

	if atomic.AddInt32(f.usage, -1) != 0 {
		return nil
	}

	listenerMut.Lock()
	defer listenerMut.Unlock()
	delete(listeners, f.key)
	logger.Infof("Close listen to %s", f.key)
	return f.Listener.Close()
}

type globalListener struct {
	usage       int32
	listener    net.Listener
	deadline    bool
	deadlineMut sync.Mutex
}

var (
	listeners   = map[string]*globalListener{}
	listenerMut sync.Mutex
)
