package destination

import (
	"net"

	"github.com/pipeproxy/pipe/components/common/register"
	"github.com/pipeproxy/pipe/components/stream"
)

const (
	name = "destination"
)

func init() {
	register.Register(name, NewDestinationWithConfig)
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

func NewDestinationWithConfig(conf *Config) (stream.Handler, error) {
	mux := NewDestination()
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
