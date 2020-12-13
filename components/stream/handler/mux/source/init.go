package source

import (
	"net"

	"github.com/pipeproxy/pipe/components/common/register"
	"github.com/pipeproxy/pipe/components/stream"
)

const (
	name = "source"
)

func init() {
	register.Register(name, NewSourceWithConfig)
}

type Route struct {
	Ports   []uint32 `json:",omitempty"`
	CIDR    string   `json:",omitempty"`
	Handler stream.Handler
}

type Config struct {
	Destinations []*Route
	NotFound     stream.Handler
}

func NewSourceWithConfig(conf *Config) (stream.Handler, error) {
	mux := NewSource()
	mux.NotFound(conf.NotFound)
	for _, destination := range conf.Destinations {
		if destination.Handler == nil {
			continue
		}
		var cidr *net.IPNet
		if destination.CIDR != "" {
			_, cidr, _ = net.ParseCIDR(destination.CIDR)
		}
		mux.Handle(destination.Ports, cidr, destination.Handler)
	}
	return mux, nil
}
