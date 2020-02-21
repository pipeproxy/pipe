package none

import (
	"context"
)

type None struct{}

func (None) Do(ctx context.Context) error {

	return nil
}
