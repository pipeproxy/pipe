package components

import "context"

type none struct {
}

func (none) Do(ctx context.Context) error {
	return nil
}
