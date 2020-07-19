// DO NOT EDIT! Code generated.
package reference

import (
	"github.com/wzshiming/pipe/components/codec"
	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/internal/logger"
)

func init() {
	register.Register("ref", NewMarshalerRefWithConfig)
	register.Register("def", NewMarshalerDefWithConfig)
	register.Register("none", NewMarshalerNone)
}

type Config struct {
	Name string
	Def  codec.Marshaler `json:",omitempty"`
}

func NewMarshalerRefWithConfig(conf *Config) (codec.Marshaler, error) {
	o := &Marshaler{
		Name: conf.Name,
		Def:  conf.Def,
	}
	return o, nil
}

func NewMarshalerDefWithConfig(conf *Config) (codec.Marshaler, error) {
	MarshalerStore[conf.Name] = conf.Def
	return conf.Def, nil
}

var MarshalerStore = map[string]codec.Marshaler{}

func MarshalerFind(name string, defaults codec.Marshaler) codec.Marshaler {
	o, ok := MarshalerStore[name]
	if ok {
		return o
	}
	if defaults != nil {
		return defaults
	}
	return MarshalerNone{}
}

type MarshalerNone struct{}

func NewMarshalerNone() codec.Marshaler {
	return MarshalerNone{}
}

func (MarshalerNone) Marshal(_ interface{}) (_ []uint8, _ error) {
	logger.Warn("this is none of codec.Marshaler")
	return
}

type Marshaler struct {
	Name string
	Def  codec.Marshaler
}

func (o *Marshaler) Marshal(a interface{}) ([]uint8, error) {
	return MarshalerFind(o.Name, o.Def).Marshal(a)
}
