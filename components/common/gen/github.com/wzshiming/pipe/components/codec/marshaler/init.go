// DO NOT EDIT! Code generated.
package marshaler

import (
	"context"
	"fmt"

	"github.com/wzshiming/pipe/components/codec"
	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/internal/ctxcache"
	"github.com/wzshiming/pipe/internal/logger"
)

func init() {
	register.Register("ref", NewMarshalerRefWithConfig)
	register.Register("def", NewMarshalerDefWithConfig)
	register.Register("none", newMarshalerNone)
}

type Config struct {
	Name string
	Def  codec.Marshaler `json:",omitempty"`
}

func NewMarshalerRefWithConfig(ctx context.Context, conf *Config) codec.Marshaler {
	o := &Marshaler{
		Name: conf.Name,
		Def:  conf.Def,
		Ctx:  ctx,
	}
	return o
}

func NewMarshalerDefWithConfig(ctx context.Context, conf *Config) codec.Marshaler {
	return MarshalerPut(ctx, conf.Name, conf.Def)
}

func MarshalerPut(ctx context.Context, name string, def codec.Marshaler) codec.Marshaler {
	if def == nil {
		def = MarshalerNone
	}

	m, ok := ctxcache.GetCacheWithContext(ctx)
	if !ok {
		return MarshalerNone
	}
	store, _ := m.LoadOrStore("codec.Marshaler", map[string]codec.Marshaler{})
	store.(map[string]codec.Marshaler)[name] = def
	return def
}

func MarshalerGet(ctx context.Context, name string, defaults codec.Marshaler) codec.Marshaler {
	m, ok := ctxcache.GetCacheWithContext(ctx)
	if ok {
		store, _ := m.LoadOrStore("codec.Marshaler", map[string]codec.Marshaler{})
		o, ok := store.(map[string]codec.Marshaler)[name]
		if ok {
			return o
		}
	}

	if defaults != nil {
		return defaults
	}
	return MarshalerNone
}

var MarshalerNone _MarshalerNone

type _MarshalerNone struct{}

func newMarshalerNone() codec.Marshaler {
	return MarshalerNone
}

func (_MarshalerNone) Marshal(_ interface{}) (_ []uint8, error error) {
	logger.Warn("this is none of codec.Marshaler")

	error = fmt.Errorf("error none")

	return
}

type Marshaler struct {
	Name string
	Def  codec.Marshaler
	Ctx  context.Context
}

func (o *Marshaler) Marshal(a interface{}) ([]uint8, error) {
	return MarshalerGet(o.Ctx, o.Name, o.Def).Marshal(a)
}
