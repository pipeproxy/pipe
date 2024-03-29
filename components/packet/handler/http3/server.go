package quic

import (
	"context"
	"net"
	"net/http"

	http3 "github.com/lucas-clemente/quic-go/http3"
	"github.com/pipeproxy/pipe/components/packet"
	"github.com/pipeproxy/pipe/components/tls"
	"github.com/pipeproxy/pipe/internal/listener"
	"github.com/wzshiming/logger"
)

type server struct {
	handler   http.Handler
	tlsConfig tls.TLS
}

func NewServer(handler http.Handler, tlsConfig tls.TLS) packet.Handler {
	s := &server{
		handler:   handler,
		tlsConfig: tlsConfig,
	}
	return s
}

func (s *server) ServePacket(ctx context.Context, pkt packet.Packet) {
	log := logger.FromContext(ctx)
	if log.Enabled() {
		log = log.WithName("http3")
		ctx = logger.WithContext(ctx, log)
	}

	httpServer := &http.Server{
		BaseContext: func(listener net.Listener) context.Context {
			return ctx
		},
		TLSConfig: s.tlsConfig.TLS(),
		Handler:   s.handler,
	}
	quicServer := &http3.Server{
		Server: httpServer,
	}
	go func() {
		<-ctx.Done()
		quicServer.Close()
	}()

	err := quicServer.Serve(pkt)
	if err != nil && !listener.IsClosedConnError(err) && !listener.IsServerClosedError(err) {
		log.Error(err, "http3 server close")
	}
}
