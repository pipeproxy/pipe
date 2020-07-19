// DO NOT EDIT! Code generated.
package reference

import (
	"io"

	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/internal/logger"
)

func init() {
	register.Register("ref", NewReaderRefWithConfig)
	register.Register("def", NewReaderDefWithConfig)
	register.Register("none", NewReaderNone)
}

type Config struct {
	Name string
	Def  io.Reader `json:",omitempty"`
}

func NewReaderRefWithConfig(conf *Config) (io.Reader, error) {
	o := &Reader{
		Name: conf.Name,
		Def:  conf.Def,
	}
	return o, nil
}

func NewReaderDefWithConfig(conf *Config) (io.Reader, error) {
	ReaderStore[conf.Name] = conf.Def
	return conf.Def, nil
}

var ReaderStore = map[string]io.Reader{}

func ReaderFind(name string, defaults io.Reader) io.Reader {
	o, ok := ReaderStore[name]
	if ok {
		return o
	}
	if defaults != nil {
		return defaults
	}
	return ReaderNone{}
}

type ReaderNone struct{}

func NewReaderNone() io.Reader {
	return ReaderNone{}
}

func (ReaderNone) Read(_ []uint8) (_ int, _ error) {
	logger.Warn("this is none of io.Reader")
	return
}

type Reader struct {
	Name string
	Def  io.Reader
}

func (o *Reader) Read(a []uint8) (int, error) {
	return ReaderFind(o.Name, o.Def).Read(a)
}
