// DO NOT EDIT! Code generated.
package reference

import (
	"io"

	"github.com/wzshiming/pipe/components/codec"
	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/internal/logger"
)

func init() {
	register.Register("ref", NewEncoderRefWithConfig)
	register.Register("def", NewEncoderDefWithConfig)
	register.Register("none", NewEncoderNone)
}

type Config struct {
	Name string
	Def  codec.Encoder `json:",omitempty"`
}

func NewEncoderRefWithConfig(conf *Config) (codec.Encoder, error) {
	o := &Encoder{
		Name: conf.Name,
		Def:  conf.Def,
	}
	return o, nil
}

func NewEncoderDefWithConfig(conf *Config) (codec.Encoder, error) {
	EncoderStore[conf.Name] = conf.Def
	return conf.Def, nil
}

var EncoderStore = map[string]codec.Encoder{}

func EncoderFind(name string, defaults codec.Encoder) codec.Encoder {
	o, ok := EncoderStore[name]
	if ok {
		return o
	}
	if defaults != nil {
		return defaults
	}
	return EncoderNone{}
}

type EncoderNone struct{}

func NewEncoderNone() codec.Encoder {
	return EncoderNone{}
}

func (EncoderNone) Encode(_ io.Writer) (_ io.Writer, _ error) {
	logger.Warn("this is none of codec.Encoder")
	return
}

type Encoder struct {
	Name string
	Def  codec.Encoder
}

func (o *Encoder) Encode(writer io.Writer) (io.Writer, error) {
	return EncoderFind(o.Name, o.Def).Encode(writer)
}
