// DO NOT EDIT! Code generated.
package reference

import (
	"net/http"

	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/internal/logger"
)

func init() {
	register.Register("ref", NewRoundTripperRefWithConfig)
	register.Register("def", NewRoundTripperDefWithConfig)
	register.Register("none", NewRoundTripperNone)
}

type Config struct {
	Name string
	Def  http.RoundTripper `json:",omitempty"`
}

func NewRoundTripperRefWithConfig(conf *Config) (http.RoundTripper, error) {
	o := &RoundTripper{
		Name: conf.Name,
		Def:  conf.Def,
	}
	return o, nil
}

func NewRoundTripperDefWithConfig(conf *Config) (http.RoundTripper, error) {
	RoundTripperStore[conf.Name] = conf.Def
	return conf.Def, nil
}

var RoundTripperStore = map[string]http.RoundTripper{}

func RoundTripperFind(name string, defaults http.RoundTripper) http.RoundTripper {
	o, ok := RoundTripperStore[name]
	if ok {
		return o
	}
	if defaults != nil {
		return defaults
	}
	return RoundTripperNone{}
}

type RoundTripperNone struct{}

func NewRoundTripperNone() http.RoundTripper {
	return RoundTripperNone{}
}

func (RoundTripperNone) RoundTrip(_ *http.Request) (_ *http.Response, _ error) {
	logger.Warn("this is none of http.RoundTripper")
	return
}

type RoundTripper struct {
	Name string
	Def  http.RoundTripper
}

func (o *RoundTripper) RoundTrip(a *http.Request) (*http.Response, error) {
	return RoundTripperFind(o.Name, o.Def).RoundTrip(a)
}
