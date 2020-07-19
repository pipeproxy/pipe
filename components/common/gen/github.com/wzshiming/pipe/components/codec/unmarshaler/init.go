// DO NOT EDIT! Code generated.
package reference

import (
	"github.com/wzshiming/pipe/components/codec"
	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/internal/logger"
)

func init() {
	register.Register("ref", NewUnmarshalerRefWithConfig)
	register.Register("def", NewUnmarshalerDefWithConfig)
	register.Register("none", NewUnmarshalerNone)
}

type Config struct {
	Name string
	Def  codec.Unmarshaler `json:",omitempty"`
}

func NewUnmarshalerRefWithConfig(conf *Config) (codec.Unmarshaler, error) {
	o := &Unmarshaler{
		Name: conf.Name,
		Def:  conf.Def,
	}
	return o, nil
}

func NewUnmarshalerDefWithConfig(conf *Config) (codec.Unmarshaler, error) {
	UnmarshalerStore[conf.Name] = conf.Def
	return conf.Def, nil
}

var UnmarshalerStore = map[string]codec.Unmarshaler{}

func UnmarshalerFind(name string, defaults codec.Unmarshaler) codec.Unmarshaler {
	o, ok := UnmarshalerStore[name]
	if ok {
		return o
	}
	if defaults != nil {
		return defaults
	}
	return UnmarshalerNone{}
}

type UnmarshalerNone struct{}

func NewUnmarshalerNone() codec.Unmarshaler {
	return UnmarshalerNone{}
}

func (UnmarshalerNone) Unmarshal(_ []uint8, _ interface{}) (_ error) {
	logger.Warn("this is none of codec.Unmarshaler")
	return
}

type Unmarshaler struct {
	Name string
	Def  codec.Unmarshaler
}

func (o *Unmarshaler) Unmarshal(a []uint8, b interface{}) error {
	return UnmarshalerFind(o.Name, o.Def).Unmarshal(a, b)
}
