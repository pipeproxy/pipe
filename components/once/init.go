package once

import (
	"context"

	"github.com/wzshiming/pipe/components/common/types"
)

func init() {
	var once Once
	types.Register(&once)
}

type Once interface {
	Do(ctx context.Context) error
}
