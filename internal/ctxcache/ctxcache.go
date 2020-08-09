package ctxcache

import (
	"context"
	"sync"
)

type cacheCtxKeyType int

func GetCacheWithContext(ctx context.Context) (*sync.Map, bool) {
	i := ctx.Value(cacheCtxKeyType(0))
	if i == nil {
		return nil, false
	}
	p, ok := i.(*sync.Map)
	return p, ok
}

func WithCache(ctx context.Context) context.Context {
	_, ok := GetCacheWithContext(ctx)
	if ok {
		return ctx
	}
	return context.WithValue(ctx, cacheCtxKeyType(0), &sync.Map{})
}
