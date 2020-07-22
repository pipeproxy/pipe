// DO NOT EDIT! Code generated.
package reference

import (
	"io"
	"sync"

	"github.com/wzshiming/pipe/components/codec"
	"github.com/wzshiming/pipe/components/common/register"
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

func NewEncoderRefWithConfig(conf *Config) codec.Encoder {
	o := &Encoder{
		Name: conf.Name,
		Def:  conf.Def,
	}
	return o
}

func NewEncoderDefWithConfig(conf *Config) codec.Encoder {
	return EncoderPut(conf.Name, conf.Def)
}

var (
	mut           sync.RWMutex
	_EncoderStore = map[string]codec.Encoder{}
)

func EncoderPut(name string, def codec.Encoder) codec.Encoder {
	if def == nil {
		def = EncoderNone
	}
	mut.Lock()
	_EncoderStore[name] = def
	mut.Unlock()
	return def
}

func EncoderGet(name string, defaults codec.Encoder) codec.Encoder {
	mut.RLock()
	o, ok := _EncoderStore[name]
	mut.RUnlock()
	if ok {
		return o
	}
	if defaults != nil {
		return defaults
	}
	return EncoderNone
}

var EncoderNone _EncoderNone

type _EncoderNone struct{}

func newEncoderNone() codec.Encoder {
	return EncoderNone
}

func (_EncoderNone) Encode(_ io.Writer) (_ io.Writer, _ error) {
	logger.Warn("this is none of codec.Encoder")
	return
}

type Encoder struct {
	Name string
	Def  codec.Encoder
}

func (o *Encoder) Encode(writer io.Writer) (io.Writer, error) {
	return EncoderGet(o.Name, o.Def).Encode(writer)
}
