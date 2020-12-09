package mux

import (
	"context"

	"github.com/wzshiming/logger"

	"github.com/pipeproxy/pipe/components/common/register"
	"github.com/pipeproxy/pipe/components/stream"
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
func NewMuxWithConfig(ctx context.Context, conf *Config) (stream.Handler, error) {
	mux := NewMux()
	mux.NotFound(conf.NotFound)

	for _, route := range conf.Routes {
		if route.Pattern != "" {
			patterm, ok := Get(route.Pattern)
			if ok && patterm != "" {
				mux.HandleRegexp(patterm, route.Handler)
			}
		} else if route.Regexp != "" {
			err := mux.HandleRegexp(route.Regexp, route.Handler)
			if err != nil {
				logger.FromContext(ctx).Error(err, "Stream mux regexp")
			}
		} else if route.Prefix != "" {
			mux.HandlePrefix(route.Prefix, route.Handler)
		}
	}
	return mux, nil
}
