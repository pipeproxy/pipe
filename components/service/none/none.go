package none

import (
	"context"
	"sync"
)

type none struct {
	ch   chan struct{}
	once sync.Once
}

func newNone() *none {
	return &none{
		ch: make(chan struct{}),
	}
}

func (n *none) Run(ctx context.Context) error {
	select {
	case <-ctx.Done():
		_ = n.Close()
	case <-n.ch:
	}
	return nil
}

func (n *none) Close() error {
	n.once.Do(func() {
		close(n.ch)
	})
	return nil
}
