package multi

import (
	"context"
	"fmt"

	"github.com/pipeproxy/pipe/components/stream"
)

var (
	ErrNotHandler = fmt.Errorf("not handler")
)

type Multi struct {
	multi []stream.Handler
}

func NewMulti(multi []stream.Handler) *Multi {
	return &Multi{
		multi: multi,
	}
}

func (m *Multi) ServeStream(ctx context.Context, stm stream.Stream) {
	switch len(m.multi) {
	case 0:
	case 1:
		handler := m.multi[0]
		handler.ServeStream(ctx, stm)
	default:
		handlers := m.multi
		for _, handler := range handlers {
			handler.ServeStream(ctx, stm)
		}
	}
}
