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
	register.Register("ref", NewReaderRefWithConfig)
	register.Register("def", NewReaderDefWithConfig)
	register.Register("none", newReaderNone)
}

type Config struct {
	Name string
	Def  io.Reader `json:",omitempty"`
}

func NewReaderRefWithConfig(conf *Config) io.Reader {
	o := &Reader{
		Name: conf.Name,
		Def:  conf.Def,
	}
	return o
}

func NewReaderDefWithConfig(conf *Config) io.Reader {
	return ReaderPut(conf.Name, conf.Def)
}

var (
	mut          sync.RWMutex
	_ReaderStore = map[string]io.Reader{}
)

func ReaderPut(name string, def io.Reader) io.Reader {
	if def == nil {
		def = ReaderNone
	}
	mut.Lock()
	_ReaderStore[name] = def
	mut.Unlock()
	return def
}

func ReaderGet(name string, defaults io.Reader) io.Reader {
	mut.RLock()
	o, ok := _ReaderStore[name]
	mut.RUnlock()
	if ok {
		return o
	}
	if defaults != nil {
		return defaults
	}
	return ReaderNone
}

var ReaderNone _ReaderNone

type _ReaderNone struct{}

func newReaderNone() io.Reader {
	return ReaderNone
}

func (_ReaderNone) Read(_ []uint8) (_ int, error error) {
	logger.Warn("this is none of io.Reader")

	error = fmt.Errorf("error none")

	return
}

type Reader struct {
	Name string
	Def  io.Reader
}

func (o *Reader) Read(a []uint8) (int, error) {
	return ReaderGet(o.Name, o.Def).Read(a)
}
