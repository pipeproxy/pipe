// DO NOT EDIT! Code generated.
package writer

import (
	"context"
	"fmt"
	"io"

	"github.com/pipeproxy/pipe/components/common/register"
	"github.com/pipeproxy/pipe/internal/ctxcache"
	"github.com/pipeproxy/pipe/internal/logger"
)

func init() {
	register.Register("ref", NewWriterRefWithConfig)
	register.Register("def", NewWriterDefWithConfig)
	register.Register("none", newWriterNone)
}

type Config struct {
	Name string
	Def  io.Writer `json:",omitempty"`
}

func NewWriterRefWithConfig(ctx context.Context, conf *Config) io.Writer {
	o := &Writer{
		Name: conf.Name,
		Def:  conf.Def,
		Ctx:  ctx,
	}
	return o
}

func NewWriterDefWithConfig(ctx context.Context, conf *Config) io.Writer {
	return WriterPut(ctx, conf.Name, conf.Def)
}

func WriterPut(ctx context.Context, name string, def io.Writer) io.Writer {
	if def == nil {
		return WriterNone
	}

	m, ok := ctxcache.GetCacheWithContext(ctx)
	if !ok {
		return WriterNone
	}
	store, _ := m.LoadOrStore("io.Writer", map[string]io.Writer{})
	store.(map[string]io.Writer)[name] = def
	return def
}

func WriterGet(ctx context.Context, name string, defaults io.Writer) io.Writer {
	m, ok := ctxcache.GetCacheWithContext(ctx)
	if ok {
		store, ok := m.Load("io.Writer")
		if ok {
			o, ok := store.(map[string]io.Writer)[name]
			if ok {
				return o
			}
		}
	}

	if defaults != nil {
		return defaults
	}
	logger.Warnf("io.Writer %q is not defined", name)
	return WriterNone
}

var WriterNone _WriterNone

type _WriterNone struct{}

func newWriterNone() io.Writer {
	return WriterNone
}

func (_WriterNone) Write(_ []uint8) (_ int, error error) {
	logger.Warn("this is none of io.Writer")

	error = fmt.Errorf("error io.Writer is none")

	return
}

type Writer struct {
	Name string
	Def  io.Writer
	Ctx  context.Context
}

func (o *Writer) Write(a []uint8) (int, error) {
	return WriterGet(o.Ctx, o.Name, o.Def).Write(a)
}
