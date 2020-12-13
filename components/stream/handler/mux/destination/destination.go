package destination

import (
	"context"
	"net"

	svc_stream "github.com/pipeproxy/pipe/components/service/stream"
	"github.com/pipeproxy/pipe/components/stream"
	"github.com/wzshiming/logger"
)

// Destination is an host multiplexer.
type Destination struct {
	routes   []route
	notFound stream.Handler
}

type route struct {
	ports   []uint32
	cidr    *net.IPNet
	handler stream.Handler
}

func NewDestination() *Destination {
	d := &Destination{}
	return d
}

func (d *Destination) NotFound(handler stream.Handler) {
	d.notFound = handler
}

func (d *Destination) Handle(ports []uint32, cidr *net.IPNet, handler stream.Handler) {
	d.routes = append(d.routes, route{
		ports:   ports,
		cidr:    cidr,
		handler: handler,
	})
}

// Handler returns route handler.
func (d *Destination) Handler(ip net.IP, port uint32) (handler stream.Handler) {
	for _, route := range d.routes {
		if route.cidr != nil {
			if !route.cidr.Contains(ip) {
				continue
			}
		}
		if len(route.ports) == 0 {
			return route.handler
		}
		for _, p := range route.ports {
			if p == port {
				return route.handler
			}
		}
	}
	return d.notFound
}

func (d *Destination) ServeStream(ctx context.Context, stm stream.Stream) {
	var dst net.Addr
	if _, addr, ok := svc_stream.GetRawStreamAndOriginalDestinationAddrWithContext(ctx); ok {
		dst = addr
	} else {
		dst = stm.LocalAddr()
	}

	var ip net.IP
	var port uint32
	if a, ok := dst.(*net.TCPAddr); ok {
		ip = a.IP
		port = uint32(a.Port)
	} else {
		host, p, err := net.SplitHostPort(dst.String())
		if err != nil {
			logger.FromContext(ctx).Error(err, "SplitHostPort")
			return
		}
		s, err := net.LookupPort(dst.Network(), p)
		if err != nil {
			logger.FromContext(ctx).Error(err, "LookupPort")
			return
		}
		ip = net.ParseIP(host)
		port = uint32(s)
	}
	handle := d.Handler(ip, port)
	if handle == nil {
		return
	}
	handle.ServeStream(ctx, stm)
}
