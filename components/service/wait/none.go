package wait

import (
	"context"
	"sync"
)

type wait struct {
	ch   chan struct{}
	once sync.Once
}

func newWait() *wait {
	return &wait{
		ch: make(chan struct{}),
	}
}

func (n *wait) Run(ctx context.Context) error {
	select {
	case <-ctx.Done():
		_ = n.Close()
	case <-n.ch:
	}
	return nil
}

func (n *wait) Close() error {
	n.once.Do(func() {
		close(n.ch)
	})
	return nil
}
