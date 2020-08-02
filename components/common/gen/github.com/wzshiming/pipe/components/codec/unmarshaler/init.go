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
	register.Register("ref", NewUnmarshalerRefWithConfig)
	register.Register("def", NewUnmarshalerDefWithConfig)
	register.Register("none", newUnmarshalerNone)
}

type Config struct {
	Name string
	Def  codec.Unmarshaler `json:",omitempty"`
}

func NewUnmarshalerRefWithConfig(conf *Config) codec.Unmarshaler {
	o := &Unmarshaler{
		Name: conf.Name,
		Def:  conf.Def,
	}
	return o
}

func NewUnmarshalerDefWithConfig(conf *Config) codec.Unmarshaler {
	return UnmarshalerPut(conf.Name, conf.Def)
}

var (
	mut               sync.RWMutex
	_UnmarshalerStore = map[string]codec.Unmarshaler{}
)

func UnmarshalerPut(name string, def codec.Unmarshaler) codec.Unmarshaler {
	if def == nil {
		def = UnmarshalerNone
	}
	mut.Lock()
	_UnmarshalerStore[name] = def
	mut.Unlock()
	return def
}

func UnmarshalerGet(name string, defaults codec.Unmarshaler) codec.Unmarshaler {
	mut.RLock()
	o, ok := _UnmarshalerStore[name]
	mut.RUnlock()
	if ok {
		return o
	}
	if defaults != nil {
		return defaults
	}
	return UnmarshalerNone
}

var UnmarshalerNone _UnmarshalerNone

type _UnmarshalerNone struct{}

func newUnmarshalerNone() codec.Unmarshaler {
	return UnmarshalerNone
}

func (_UnmarshalerNone) Unmarshal(_ []uint8, _ interface{}) (error error) {
	logger.Warn("this is none of codec.Unmarshaler")

	error = fmt.Errorf("error none")

	return
}

type Unmarshaler struct {
	Name string
	Def  codec.Unmarshaler
}

func (o *Unmarshaler) Unmarshal(a []uint8, b interface{}) error {
	return UnmarshalerGet(o.Name, o.Def).Unmarshal(a, b)
}
