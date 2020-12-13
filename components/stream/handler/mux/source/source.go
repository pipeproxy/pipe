package source

import (
	"context"
	"net"

	"github.com/pipeproxy/pipe/components/stream"
	"github.com/wzshiming/logger"
)

// Source is an host multiplexer.
type Source struct {
	routes   []route
	notFound stream.Handler
}

type route struct {
	ports   []uint32
	cidr    *net.IPNet
	handler stream.Handler
}

func NewSource() *Source {
	s := &Source{}
	return s
}

func (s *Source) NotFound(handler stream.Handler) {
	s.notFound = handler
}

func (s *Source) Handle(ports []uint32, cidr *net.IPNet, handler stream.Handler) {
	s.routes = append(s.routes, route{
		ports:   ports,
		cidr:    cidr,
		handler: handler,
	})
}

// Handler returns route handler.
func (s *Source) Handler(ip net.IP, port uint32) (handler stream.Handler) {
	for _, route := range s.routes {
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
	return s.notFound
}

func (s *Source) ServeStream(ctx context.Context, stm stream.Stream) {
	src := stm.RemoteAddr()

	var ip net.IP
	var port uint32
	if a, ok := src.(*net.TCPAddr); ok {
		ip = a.IP
		port = uint32(a.Port)
	} else {
		host, p, err := net.SplitHostPort(src.String())
		if err != nil {
			logger.FromContext(ctx).Error(err, "SplitHostPort")
			return
		}
		s, err := net.LookupPort(src.Network(), p)
		if err != nil {
			logger.FromContext(ctx).Error(err, "LookupPort")
			return
		}
		ip = net.ParseIP(host)
		port = uint32(s)
	}
	handle := s.Handler(ip, port)
	if handle == nil {
		return
	}
	handle.ServeStream(ctx, stm)
}
