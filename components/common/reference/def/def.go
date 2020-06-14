package def

import (
	"context"
	"reflect"

	"github.com/wzshiming/funcfg/define"
	"github.com/wzshiming/pipe/internal/refctx"
)

func Def(ctx context.Context, name string, def define.Self, i interface{}) error {
	if def == nil {
		return refctx.ErrDefEmpty
	}
	val, ok := refctx.Get(ctx)
	if !ok {
		return refctx.ErrNotUse
	}

	s := reflect.ValueOf(i).Elem()
	d := reflect.ValueOf(def)
	s.Set(d)
	return val.Def(name, s, d)
}
