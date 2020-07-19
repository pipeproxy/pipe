// DO NOT EDIT! Code generated.
package reference

import (
	"io"

	"github.com/wzshiming/pipe/components/codec"
	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/internal/logger"
)

func init() {
	register.Register("ref", NewDecoderRefWithConfig)
	register.Register("def", NewDecoderDefWithConfig)
	register.Register("none", NewDecoderNone)
}

type Config struct {
	Name string
	Def  codec.Decoder `json:",omitempty"`
}

func NewDecoderRefWithConfig(conf *Config) (codec.Decoder, error) {
	o := &Decoder{
		Name: conf.Name,
		Def:  conf.Def,
	}
	return o, nil
}

func NewDecoderDefWithConfig(conf *Config) (codec.Decoder, error) {
	DecoderStore[conf.Name] = conf.Def
	return conf.Def, nil
}

var DecoderStore = map[string]codec.Decoder{}

func DecoderFind(name string, defaults codec.Decoder) codec.Decoder {
	o, ok := DecoderStore[name]
	if ok {
		return o
	}
	if defaults != nil {
		return defaults
	}
	return DecoderNone{}
}

type DecoderNone struct{}

func NewDecoderNone() codec.Decoder {
	return DecoderNone{}
}

func (DecoderNone) Decode(_ io.Reader) (_ io.Reader, _ error) {
	logger.Warn("this is none of codec.Decoder")
	return
}

type Decoder struct {
	Name string
	Def  codec.Decoder
}

func (o *Decoder) Decode(reader io.Reader) (io.Reader, error) {
	return DecoderFind(o.Name, o.Def).Decode(reader)
}
