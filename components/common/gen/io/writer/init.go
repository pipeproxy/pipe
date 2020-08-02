// DO NOT EDIT! Code generated.
package reference

import (
	"fmt"
	"io"
	"sync"

	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/internal/logger"
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

func NewWriterRefWithConfig(conf *Config) io.Writer {
	o := &Writer{
		Name: conf.Name,
		Def:  conf.Def,
	}
	return o
}

func NewWriterDefWithConfig(conf *Config) io.Writer {
	return WriterPut(conf.Name, conf.Def)
}

var (
	mut          sync.RWMutex
	_WriterStore = map[string]io.Writer{}
)

func WriterPut(name string, def io.Writer) io.Writer {
	if def == nil {
		def = WriterNone
	}
	mut.Lock()
	_WriterStore[name] = def
	mut.Unlock()
	return def
}

func WriterGet(name string, defaults io.Writer) io.Writer {
	mut.RLock()
	o, ok := _WriterStore[name]
	mut.RUnlock()
	if ok {
		return o
	}
	if defaults != nil {
		return defaults
	}
	return WriterNone
}

var WriterNone _WriterNone

type _WriterNone struct{}

func newWriterNone() io.Writer {
	return WriterNone
}

func (_WriterNone) Write(_ []uint8) (_ int, error error) {
	logger.Warn("this is none of io.Writer")

	error = fmt.Errorf("error none")

	return
}

type Writer struct {
	Name string
	Def  io.Writer
}

func (o *Writer) Write(a []uint8) (int, error) {
	return WriterGet(o.Name, o.Def).Write(a)
}
