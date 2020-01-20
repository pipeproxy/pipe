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
	multi []stream.Handler
}

func NewMulti(multi []stream.Handler) *Multi {
	return &Multi{
		multi: multi,
	}
}

func (m *Multi) ServeStream(ctx context.Context, stm stream.Stream) {
	handlers := m.multi
	for _, handler := range handlers {
		handler.ServeStream(ctx, stm)
	}
}
