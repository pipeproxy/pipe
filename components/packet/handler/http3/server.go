package quic

import (
	"context"
	"net"
	"net/http"

	http3 "github.com/lucas-clemente/quic-go/http3"
	"github.com/wzshiming/pipe/components/packet"
	"github.com/wzshiming/pipe/components/tls"
	"github.com/wzshiming/pipe/internal/logger"
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
	if err != nil {
		if err.Error() != "server closed" {
			logger.Errorln("[http3]", err)
		}
	}
}
