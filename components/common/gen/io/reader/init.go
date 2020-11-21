// DO NOT EDIT! Code generated.
package reader

import (
	"context"
	"fmt"
	"io"

	"github.com/pipeproxy/pipe/components/common/register"
	"github.com/pipeproxy/pipe/internal/ctxcache"
	"github.com/wzshiming/logger"
)

func init() {
	register.Register("ref", NewReaderRefWithConfig)
	register.Register("def", NewReaderDefWithConfig)
	register.Register("none", newReaderNone)
}

type Config struct {
	Name string
	Def  io.Reader `json:",omitempty"`
}

func NewReaderRefWithConfig(ctx context.Context, conf *Config) io.Reader {
	o := &Reader{
		Name: conf.Name,
		Def:  conf.Def,
		Ctx:  ctx,
	}
	return o
}

func NewReaderDefWithConfig(ctx context.Context, conf *Config) io.Reader {
	return ReaderPut(ctx, conf.Name, conf.Def)
}

func ReaderPut(ctx context.Context, name string, def io.Reader) io.Reader {
	if def == nil {
		return ReaderNone
	}

	m, ok := ctxcache.GetCacheWithContext(ctx)
	if !ok {
		return ReaderNone
	}
	store, _ := m.LoadOrStore("io.Reader", map[string]io.Reader{})
	store.(map[string]io.Reader)[name] = def
	return def
}

func ReaderGet(ctx context.Context, name string, defaults io.Reader) io.Reader {
	m, ok := ctxcache.GetCacheWithContext(ctx)
	if ok {
		store, ok := m.Load("io.Reader")
		if ok {
			o, ok := store.(map[string]io.Reader)[name]
			if ok {
				return o
			}
		}
	}

	if defaults != nil {
		return defaults
	}
	logger.FromContext(ctx).V(-1).Info("io.Reader is not defined", "name", name)
	return ReaderNone
}

var ReaderNone _ReaderNone

type _ReaderNone struct{}

func newReaderNone() io.Reader {
	return ReaderNone
}

func (_ReaderNone) Read(_ []uint8) (_ int, error error) {
	logger.Log.V(-1).Info("this is none of io.Reader")

	error = fmt.Errorf("error io.Reader is none")

	return
}

type Reader struct {
	Name string
	Def  io.Reader
	Ctx  context.Context
}

func (o *Reader) Read(a []uint8) (int, error) {
	return ReaderGet(o.Ctx, o.Name, o.Def).Read(a)
}
