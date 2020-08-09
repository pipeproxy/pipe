package listener

import (
	"context"
	"io"
	"sync"
)

var (
	closers     []io.Closer
	closersSwap []io.Closer
	closersMut  sync.Mutex
)

func channelClose(ctx context.Context, closer io.Closer) {
	closersMut.Lock()
	defer closersMut.Unlock()
	closers = append(closers, closer)
}

func Swap() {
	closersMut.Lock()
	defer closersMut.Unlock()
	closers, closersSwap = closersSwap, closers
}

func CloseSwap() {
	closersMut.Lock()
	defer closersMut.Unlock()
	for _, closer := range closersSwap {
		closer.Close()
	}
	closersSwap = closersSwap[:0]
}
