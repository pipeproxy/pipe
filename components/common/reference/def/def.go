package def

import (
	"context"
	"reflect"

	"github.com/wzshiming/funcfg/define"
	"github.com/wzshiming/pipe/components/common/reference/ctxreference"
)

func Def(ctx context.Context, name string, def define.Self, i interface{}) error {
	if def == nil {
		return ctxreference.ErrDefEmpty
	}
	val, ok := ctxreference.Get(ctx)
	if !ok {
		return ctxreference.ErrNotUse
	}

	s := reflect.ValueOf(i).Elem()
	d := reflect.ValueOf(def)
	s.Set(d)
	return val.Def(name, s, d)
}
