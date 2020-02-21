package multi

import (
	"context"
	"fmt"

	"github.com/wzshiming/pipe/pipe/once"
)

var (
	ErrNotServer = fmt.Errorf("not server")
)

type Multi struct {
	multi []once.Once
}

func NewMulti(multi []once.Once) *Multi {
	return &Multi{
		multi: multi,
	}
}

func (m *Multi) Do(ctx context.Context) error {
	for _, do := range m.multi {
		err := do.Do(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}
