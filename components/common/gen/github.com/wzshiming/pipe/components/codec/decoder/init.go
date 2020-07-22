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
	register.Register("ref", NewDecoderRefWithConfig)
	register.Register("def", NewDecoderDefWithConfig)
	register.Register("none", newDecoderNone)
}

type Config struct {
	Name string
	Def  codec.Decoder `json:",omitempty"`
}

func NewDecoderRefWithConfig(conf *Config) codec.Decoder {
	o := &Decoder{
		Name: conf.Name,
		Def:  conf.Def,
	}
	return o
}

func NewDecoderDefWithConfig(conf *Config) codec.Decoder {
	return DecoderPut(conf.Name, conf.Def)
}

var (
	mut           sync.RWMutex
	_DecoderStore = map[string]codec.Decoder{}
)

func DecoderPut(name string, def codec.Decoder) codec.Decoder {
	if def == nil {
		def = DecoderNone
	}
	mut.Lock()
	_DecoderStore[name] = def
	mut.Unlock()
	return def
}

func DecoderGet(name string, defaults codec.Decoder) codec.Decoder {
	mut.RLock()
	o, ok := _DecoderStore[name]
	mut.RUnlock()
	if ok {
		return o
	}
	if defaults != nil {
		return defaults
	}
	return DecoderNone
}

var DecoderNone _DecoderNone

type _DecoderNone struct{}

func newDecoderNone() codec.Decoder {
	return DecoderNone
}

func (_DecoderNone) Decode(_ io.Reader) (_ io.Reader, _ error) {
	logger.Warn("this is none of codec.Decoder")
	return
}

type Decoder struct {
	Name string
	Def  codec.Decoder
}

func (o *Decoder) Decode(reader io.Reader) (io.Reader, error) {
	return DecoderGet(o.Name, o.Def).Decode(reader)
}
