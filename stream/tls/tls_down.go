package tls

import (
	"context"
	"crypto/tls"

	"github.com/wzshiming/pipe/stream"
)

type TlsServer struct {
	handler   stream.Handler
	tlsConfig *tls.Config
}

func NewTlsDown(handler stream.Handler, tlsConfig *tls.Config) *TlsServer {
	return &TlsServer{
		handler:   handler,
		tlsConfig: tlsConfig,
	}
}

func (t *TlsServer) ServeStream(ctx context.Context, stm stream.Stream) {
	conn := tls.Server(stm, t.tlsConfig)
	t.handler.ServeStream(ctx, conn)
}
