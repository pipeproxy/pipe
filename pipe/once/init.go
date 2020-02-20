package once

import (
	"context"

	"github.com/wzshiming/pipe/configure/alias"
	"github.com/wzshiming/pipe/pipe/common/load"
)

func init() {
	var once Once
	alias.Register("Once", &once)
	load.Register(&once)
}

type Once interface {
	Do(ctx context.Context) error
}
