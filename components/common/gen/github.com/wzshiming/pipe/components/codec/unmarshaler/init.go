// DO NOT EDIT! Code generated.
package unmarshaler

import (
	"context"
	"fmt"

	"github.com/wzshiming/pipe/components/codec"
	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/internal/ctxcache"
	"github.com/wzshiming/pipe/internal/logger"
)

func init() {
	register.Register("ref", NewUnmarshalerRefWithConfig)
	register.Register("def", NewUnmarshalerDefWithConfig)
	register.Register("none", newUnmarshalerNone)
}

type Config struct {
	Name string
	Def  codec.Unmarshaler `json:",omitempty"`
}

func NewUnmarshalerRefWithConfig(ctx context.Context, conf *Config) codec.Unmarshaler {
	o := &Unmarshaler{
		Name: conf.Name,
		Def:  conf.Def,
		Ctx:  ctx,
	}
	return o
}

func NewUnmarshalerDefWithConfig(ctx context.Context, conf *Config) codec.Unmarshaler {
	return UnmarshalerPut(ctx, conf.Name, conf.Def)
}

func UnmarshalerPut(ctx context.Context, name string, def codec.Unmarshaler) codec.Unmarshaler {
	if def == nil {
		return UnmarshalerNone
	}

	m, ok := ctxcache.GetCacheWithContext(ctx)
	if !ok {
		return UnmarshalerNone
	}
	store, _ := m.LoadOrStore("codec.Unmarshaler", map[string]codec.Unmarshaler{})
	store.(map[string]codec.Unmarshaler)[name] = def
	return def
}

func UnmarshalerGet(ctx context.Context, name string, defaults codec.Unmarshaler) codec.Unmarshaler {
	m, ok := ctxcache.GetCacheWithContext(ctx)
	if ok {
		store, ok := m.Load("codec.Unmarshaler")
		if ok {
			o, ok := store.(map[string]codec.Unmarshaler)[name]
			if ok {
				return o
			}
		}
	}

	if defaults != nil {
		return defaults
	}
	logger.Warnf("codec.Unmarshaler %q is not defined", name)
	return UnmarshalerNone
}

var UnmarshalerNone _UnmarshalerNone

type _UnmarshalerNone struct{}

func newUnmarshalerNone() codec.Unmarshaler {
	return UnmarshalerNone
}

func (_UnmarshalerNone) Unmarshal(_ []uint8, _ interface{}) (error error) {
	logger.Warn("this is none of codec.Unmarshaler")

	error = fmt.Errorf("error codec.Unmarshaler is none")

	return
}

type Unmarshaler struct {
	Name string
	Def  codec.Unmarshaler
	Ctx  context.Context
}

func (o *Unmarshaler) Unmarshal(a []uint8, b interface{}) error {
	return UnmarshalerGet(o.Ctx, o.Name, o.Def).Unmarshal(a, b)
}
