package prefix

import (
	"context"

	"github.com/pipeproxy/pipe/components/common/register"
	"github.com/pipeproxy/pipe/components/stream"
	"github.com/wzshiming/logger"
)

const (
	name = "prefix"
)

func init() {
	register.Register(name, NewPrefixWithConfig)
}

type ProtocolEnum string

const (
	ProtocolTLS    ProtocolEnum = "tls"
	ProtocolSocks4 ProtocolEnum = "socks4"
	ProtocolSocks5 ProtocolEnum = "socks5"
	ProtocolHTTP1  ProtocolEnum = "http1"
	ProtocolHTTP2  ProtocolEnum = "http2"
	ProtocolSSH    ProtocolEnum = "ssh"
)

type Route struct {
	Pattern ProtocolEnum `json:",omitempty"`
	Regexp  string       `json:",omitempty"`
	Prefix  string       `json:",omitempty"`
	Handler stream.Handler
}

type Config struct {
	Routes   []*Route
	NotFound stream.Handler
}

// NewPrefixWithConfig create a new Prefix with config.
func NewPrefixWithConfig(ctx context.Context, conf *Config) (stream.Handler, error) {
	mux := NewPrefix()
	mux.NotFound(conf.NotFound)

	for _, route := range conf.Routes {
		if route.Pattern != "" {
			pattern, ok := Get(string(route.Pattern))
			if ok && pattern != "" {
				err := mux.HandleRegexp(pattern, route.Handler)
				if err != nil {
					logger.FromContext(ctx).Error(err, "Stream mux pattern")
				}
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
