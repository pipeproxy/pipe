// DO NOT EDIT! Code generated.
package reference

import (
	"fmt"
	"sync"

	"github.com/wzshiming/pipe/components/codec"
	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/internal/logger"
)

func init() {
	register.Register("ref", NewMarshalerRefWithConfig)
	register.Register("def", NewMarshalerDefWithConfig)
	register.Register("none", newMarshalerNone)
}

type Config struct {
	Name string
	Def  codec.Marshaler `json:",omitempty"`
}

func NewMarshalerRefWithConfig(conf *Config) codec.Marshaler {
	o := &Marshaler{
		Name: conf.Name,
		Def:  conf.Def,
	}
	return o
}

func NewMarshalerDefWithConfig(conf *Config) codec.Marshaler {
	return MarshalerPut(conf.Name, conf.Def)
}

var (
	mut             sync.RWMutex
	_MarshalerStore = map[string]codec.Marshaler{}
)

func MarshalerPut(name string, def codec.Marshaler) codec.Marshaler {
	if def == nil {
		def = MarshalerNone
	}
	mut.Lock()
	_MarshalerStore[name] = def
	mut.Unlock()
	return def
}

func MarshalerGet(name string, defaults codec.Marshaler) codec.Marshaler {
	mut.RLock()
	o, ok := _MarshalerStore[name]
	mut.RUnlock()
	if ok {
		return o
	}
	if defaults != nil {
		return defaults
	}
	return MarshalerNone
}

var MarshalerNone _MarshalerNone

type _MarshalerNone struct{}

func newMarshalerNone() codec.Marshaler {
	return MarshalerNone
}

func (_MarshalerNone) Marshal(_ interface{}) (_ []uint8, error error) {
	logger.Warn("this is none of codec.Marshaler")

	error = fmt.Errorf("error none")

	return
}

type Marshaler struct {
	Name string
	Def  codec.Marshaler
}

func (o *Marshaler) Marshal(a interface{}) ([]uint8, error) {
	return MarshalerGet(o.Name, o.Def).Marshal(a)
}
