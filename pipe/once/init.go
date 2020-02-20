package once

import (
	"context"

	"github.com/wzshiming/pipe/configure/alias"
)

func init() {
	var once Once
	alias.Register("Once", &once)
}

type Once interface {
	Do(ctx context.Context) error
}
