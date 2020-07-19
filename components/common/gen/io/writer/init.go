// DO NOT EDIT! Code generated.
package reference

import (
	"io"

	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/internal/logger"
)

func init() {
	register.Register("ref", NewWriterRefWithConfig)
	register.Register("def", NewWriterDefWithConfig)
	register.Register("none", NewWriterNone)
}

type Config struct {
	Name string
	Def  io.Writer `json:",omitempty"`
}

func NewWriterRefWithConfig(conf *Config) (io.Writer, error) {
	o := &Writer{
		Name: conf.Name,
		Def:  conf.Def,
	}
	return o, nil
}

func NewWriterDefWithConfig(conf *Config) (io.Writer, error) {
	WriterStore[conf.Name] = conf.Def
	return conf.Def, nil
}

var WriterStore = map[string]io.Writer{}

func WriterFind(name string, defaults io.Writer) io.Writer {
	o, ok := WriterStore[name]
	if ok {
		return o
	}
	if defaults != nil {
		return defaults
	}
	return WriterNone{}
}

type WriterNone struct{}

func NewWriterNone() io.Writer {
	return WriterNone{}
}

func (WriterNone) Write(_ []uint8) (_ int, _ error) {
	logger.Warn("this is none of io.Writer")
	return
}

type Writer struct {
	Name string
	Def  io.Writer
}

func (o *Writer) Write(a []uint8) (int, error) {
	return WriterFind(o.Name, o.Def).Write(a)
}
