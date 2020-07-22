// DO NOT EDIT! Code generated.
package reference

import (
	"net/http"
	"sync"

	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/internal/logger"
)

func init() {
	register.Register("ref", NewRoundTripperRefWithConfig)
	register.Register("def", NewRoundTripperDefWithConfig)
	register.Register("none", newRoundTripperNone)
}

type Config struct {
	Name string
	Def  http.RoundTripper `json:",omitempty"`
}

func NewRoundTripperRefWithConfig(conf *Config) http.RoundTripper {
	o := &RoundTripper{
		Name: conf.Name,
		Def:  conf.Def,
	}
	return o
}

func NewRoundTripperDefWithConfig(conf *Config) http.RoundTripper {
	return RoundTripperPut(conf.Name, conf.Def)
}

var (
	mut                sync.RWMutex
	_RoundTripperStore = map[string]http.RoundTripper{}
)

func RoundTripperPut(name string, def http.RoundTripper) http.RoundTripper {
	if def == nil {
		def = RoundTripperNone
	}
	mut.Lock()
	_RoundTripperStore[name] = def
	mut.Unlock()
	return def
}

func RoundTripperGet(name string, defaults http.RoundTripper) http.RoundTripper {
	mut.RLock()
	o, ok := _RoundTripperStore[name]
	mut.RUnlock()
	if ok {
		return o
	}
	if defaults != nil {
		return defaults
	}
	return RoundTripperNone
}

var RoundTripperNone _RoundTripperNone

type _RoundTripperNone struct{}

func newRoundTripperNone() http.RoundTripper {
	return RoundTripperNone
}

func (_RoundTripperNone) RoundTrip(_ *http.Request) (_ *http.Response, _ error) {
	logger.Warn("this is none of http.RoundTripper")
	return
}

type RoundTripper struct {
	Name string
	Def  http.RoundTripper
}

func (o *RoundTripper) RoundTrip(a *http.Request) (*http.Response, error) {
	return RoundTripperGet(o.Name, o.Def).RoundTrip(a)
}
