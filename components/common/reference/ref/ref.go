package ref

import (
	"context"
	"reflect"

	"github.com/wzshiming/pipe/internal/refctx"
)

func Ref(ctx context.Context, name string, i interface{}) error {
	val, ok := refctx.Get(ctx)
	if !ok {
		return refctx.ErrNotUse
	}
	s := reflect.ValueOf(i).Elem()
	return val.Ref(name, s)
}
