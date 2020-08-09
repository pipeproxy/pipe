// DO NOT EDIT! Code generated.
package roundtripper

import (
	"context"
	"fmt"
	"net/http"

	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/internal/ctxcache"
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

func NewRoundTripperRefWithConfig(ctx context.Context, conf *Config) http.RoundTripper {
	o := &RoundTripper{
		Name: conf.Name,
		Def:  conf.Def,
		Ctx:  ctx,
	}
	return o
}

func NewRoundTripperDefWithConfig(ctx context.Context, conf *Config) http.RoundTripper {
	return RoundTripperPut(ctx, conf.Name, conf.Def)
}

func RoundTripperPut(ctx context.Context, name string, def http.RoundTripper) http.RoundTripper {
	if def == nil {
		def = RoundTripperNone
	}

	m, ok := ctxcache.GetCacheWithContext(ctx)
	if !ok {
		return RoundTripperNone
	}
	store, _ := m.LoadOrStore("http.RoundTripper", map[string]http.RoundTripper{})
	store.(map[string]http.RoundTripper)[name] = def
	return def
}

func RoundTripperGet(ctx context.Context, name string, defaults http.RoundTripper) http.RoundTripper {
	m, ok := ctxcache.GetCacheWithContext(ctx)
	if ok {
		store, _ := m.LoadOrStore("http.RoundTripper", map[string]http.RoundTripper{})
		o, ok := store.(map[string]http.RoundTripper)[name]
		if ok {
			return o
		}
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

func (_RoundTripperNone) RoundTrip(_ *http.Request) (_ *http.Response, error error) {
	logger.Warn("this is none of http.RoundTripper")

	error = fmt.Errorf("error none")

	return
}

type RoundTripper struct {
	Name string
	Def  http.RoundTripper
	Ctx  context.Context
}

func (o *RoundTripper) RoundTrip(a *http.Request) (*http.Response, error) {
	return RoundTripperGet(o.Ctx, o.Name, o.Def).RoundTrip(a)
}
