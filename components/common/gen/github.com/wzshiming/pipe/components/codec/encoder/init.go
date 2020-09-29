// DO NOT EDIT! Code generated.
package encoder

import (
	"context"
	"fmt"
	"io"

	"github.com/wzshiming/pipe/components/codec"
	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/internal/ctxcache"
	"github.com/wzshiming/pipe/internal/logger"
)

func init() {
	register.Register("ref", NewEncoderRefWithConfig)
	register.Register("def", NewEncoderDefWithConfig)
	register.Register("none", newEncoderNone)
}

type Config struct {
	Name string
	Def  codec.Encoder `json:",omitempty"`
}

func NewEncoderRefWithConfig(ctx context.Context, conf *Config) codec.Encoder {
	o := &Encoder{
		Name: conf.Name,
		Def:  conf.Def,
		Ctx:  ctx,
	}
	return o
}

func NewEncoderDefWithConfig(ctx context.Context, conf *Config) codec.Encoder {
	return EncoderPut(ctx, conf.Name, conf.Def)
}

func EncoderPut(ctx context.Context, name string, def codec.Encoder) codec.Encoder {
	if def == nil {
		return EncoderNone
	}

	m, ok := ctxcache.GetCacheWithContext(ctx)
	if !ok {
		return EncoderNone
	}
	store, _ := m.LoadOrStore("codec.Encoder", map[string]codec.Encoder{})
	store.(map[string]codec.Encoder)[name] = def
	return def
}

func EncoderGet(ctx context.Context, name string, defaults codec.Encoder) codec.Encoder {
	m, ok := ctxcache.GetCacheWithContext(ctx)
	if ok {
		store, ok := m.Load("codec.Encoder")
		if ok {
			o, ok := store.(map[string]codec.Encoder)[name]
			if ok {
				return o
			}
		}
	}

	if defaults != nil {
		return defaults
	}
	logger.Warnf("codec.Encoder %q is not defined", name)
	return EncoderNone
}

var EncoderNone _EncoderNone

type _EncoderNone struct{}

func newEncoderNone() codec.Encoder {
	return EncoderNone
}

func (_EncoderNone) Encode(_ io.Writer) (_ io.Writer, error error) {
	logger.Warn("this is none of codec.Encoder")

	error = fmt.Errorf("error codec.Encoder is none")

	return
}

type Encoder struct {
	Name string
	Def  codec.Encoder
	Ctx  context.Context
}

func (o *Encoder) Encode(writer io.Writer) (io.Writer, error) {
	return EncoderGet(o.Ctx, o.Name, o.Def).Encode(writer)
}
