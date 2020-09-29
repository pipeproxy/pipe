// DO NOT EDIT! Code generated.
package decoder

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
	register.Register("ref", NewDecoderRefWithConfig)
	register.Register("def", NewDecoderDefWithConfig)
	register.Register("none", newDecoderNone)
}

type Config struct {
	Name string
	Def  codec.Decoder `json:",omitempty"`
}

func NewDecoderRefWithConfig(ctx context.Context, conf *Config) codec.Decoder {
	o := &Decoder{
		Name: conf.Name,
		Def:  conf.Def,
		Ctx:  ctx,
	}
	return o
}

func NewDecoderDefWithConfig(ctx context.Context, conf *Config) codec.Decoder {
	return DecoderPut(ctx, conf.Name, conf.Def)
}

func DecoderPut(ctx context.Context, name string, def codec.Decoder) codec.Decoder {
	if def == nil {
		return DecoderNone
	}

	m, ok := ctxcache.GetCacheWithContext(ctx)
	if !ok {
		return DecoderNone
	}
	store, _ := m.LoadOrStore("codec.Decoder", map[string]codec.Decoder{})
	store.(map[string]codec.Decoder)[name] = def
	return def
}

func DecoderGet(ctx context.Context, name string, defaults codec.Decoder) codec.Decoder {
	m, ok := ctxcache.GetCacheWithContext(ctx)
	if ok {
		store, ok := m.Load("codec.Decoder")
		if ok {
			o, ok := store.(map[string]codec.Decoder)[name]
			if ok {
				return o
			}
		}
	}

	if defaults != nil {
		return defaults
	}
	logger.Warnf("codec.Decoder %q is not defined", name)
	return DecoderNone
}

var DecoderNone _DecoderNone

type _DecoderNone struct{}

func newDecoderNone() codec.Decoder {
	return DecoderNone
}

func (_DecoderNone) Decode(_ io.Reader) (_ io.Reader, error error) {
	logger.Warn("this is none of codec.Decoder")

	error = fmt.Errorf("error codec.Decoder is none")

	return
}

type Decoder struct {
	Name string
	Def  codec.Decoder
	Ctx  context.Context
}

func (o *Decoder) Decode(reader io.Reader) (io.Reader, error) {
	return DecoderGet(o.Ctx, o.Name, o.Def).Decode(reader)
}
