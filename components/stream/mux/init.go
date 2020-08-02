package mux

import (
	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/components/stream"
)

const (
	name = "mux"
)

func init() {
	register.Register(name, NewMuxWithConfig)
}

type Route struct {
	Pattern string `json:",omitempty"`
	Regexp  string `json:",omitempty"`
	Prefix  string `json:",omitempty"`
	Handler stream.Handler
}

type Config struct {
	Routes   []*Route
	NotFound stream.Handler
}

// NewProtoMux create a new Mux with config.
func NewMuxWithConfig(conf *Config) (stream.Handler, error) {
	mux := NewMux()
	if conf.NotFound != nil {
		mux.NotFound(conf.NotFound)
	}

	for _, route := range conf.Routes {
		if route.Pattern != "" {
			patterm, ok := Get(route.Pattern)
			if ok && patterm != "" {
				mux.HandleRegexp(patterm, route.Handler)
			}
		} else if route.Regexp != "" {
			mux.HandleRegexp(route.Regexp, route.Handler)
		} else if route.Prefix != "" {
			mux.HandlePrefix(route.Prefix, route.Handler)
		}
	}
	return mux, nil
}
