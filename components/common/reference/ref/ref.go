package ref

import (
	"context"
	"reflect"

	"github.com/wzshiming/pipe/components/common/reference/ctxreference"
)

func Ref(ctx context.Context, name string, i interface{}) error {
	val, ok := ctxreference.Get(ctx)
	if !ok {
		return ctxreference.ErrNotUse
	}
	s := reflect.ValueOf(i).Elem()
	return val.Ref(name, s)
}
