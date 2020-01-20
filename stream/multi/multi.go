package multi

import (
	"context"
	"fmt"

	"github.com/wzshiming/pipe/stream"
)

var (
	ErrNotHandler = fmt.Errorf("not handler")
)

type Multi struct {
	handlers []stream.Handler
}

func NewMulti(handlers []stream.Handler) *Multi {
	return &Multi{
		handlers: handlers,
	}
}

func (m *Multi) ServeStream(ctx context.Context, stm stream.Stream) {
	handlers := m.handlers
	for _, handler := range handlers {
		handler.ServeStream(ctx, stm)
	}
}
