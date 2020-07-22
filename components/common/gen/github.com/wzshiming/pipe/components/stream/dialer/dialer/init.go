// DO NOT EDIT! Code generated.
package reference

import (
	"context"
	"net"
	"sync"

	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/components/stream/dialer"
	"github.com/wzshiming/pipe/internal/logger"
)

func init() {
	register.Register("ref", NewDialerRefWithConfig)
	register.Register("def", NewDialerDefWithConfig)
	register.Register("none", newDialerNone)
}

type Config struct {
	Name string
	Def  dialer.Dialer `json:",omitempty"`
}

func NewDialerRefWithConfig(conf *Config) dialer.Dialer {
	o := &Dialer{
		Name: conf.Name,
		Def:  conf.Def,
	}
	return o
}

func NewDialerDefWithConfig(conf *Config) dialer.Dialer {
	return DialerPut(conf.Name, conf.Def)
}

var (
	mut          sync.RWMutex
	_DialerStore = map[string]dialer.Dialer{}
)

func DialerPut(name string, def dialer.Dialer) dialer.Dialer {
	if def == nil {
		def = DialerNone
	}
	mut.Lock()
	_DialerStore[name] = def
	mut.Unlock()
	return def
}

func DialerGet(name string, defaults dialer.Dialer) dialer.Dialer {
	mut.RLock()
	o, ok := _DialerStore[name]
	mut.RUnlock()
	if ok {
		return o
	}
	if defaults != nil {
		return defaults
	}
	return DialerNone
}

var DialerNone _DialerNone

type _DialerNone struct{}

func newDialerNone() dialer.Dialer {
	return DialerNone
}

func (_DialerNone) DialStream(_ context.Context) (_ net.Conn, _ error) {
	logger.Warn("this is none of dialer.Dialer")
	return
}

type Dialer struct {
	Name string
	Def  dialer.Dialer
}

func (o *Dialer) DialStream(context context.Context) (net.Conn, error) {
	return DialerGet(o.Name, o.Def).DialStream(context)
}
