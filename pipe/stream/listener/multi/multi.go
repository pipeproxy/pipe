package multi

import (
	"context"
	"fmt"
	"net"

	"github.com/wzshiming/pipe/pipe/stream/listener"
)

var (
	ErrNotListener = fmt.Errorf("not listener")
)

type Multi struct {
	multi []listener.ListenConfig
}

func NewMulti(multi []listener.ListenConfig) *Multi {
	return &Multi{
		multi: multi,
	}
}

func (m *Multi) ListenStream(ctx context.Context) (net.Listener, error) {
	multi := make([]net.Listener, 0, len(m.multi))
	for _, lc := range m.multi {
		l, err := lc.ListenStream(ctx)
		if err != nil {
			fmt.Errorf("[ERROR] listen")
			continue
		}
		multi = append(multi, l)
	}
	switch len(multi) {
	case 0:
		return nil, ErrNotListener
	case 1:
		return multi[0], nil
	}
	return newMultiListener(multi)
}
