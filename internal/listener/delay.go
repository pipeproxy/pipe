package listener

import (
	"context"
	"io"
	"sync"
)

var (
	closersSwap = map[string]io.Closer{}
	closersMut  sync.Mutex
)

type keyCloser interface {
	Key() string
	io.Closer
}

func swapClose(ctx context.Context, closer keyCloser) {
	closersMut.Lock()
	defer closersMut.Unlock()
	key := closer.Key()
	c, ok := closersSwap[key]
	if ok {
		c.Close()
	}
	closersSwap[key] = closer
}
