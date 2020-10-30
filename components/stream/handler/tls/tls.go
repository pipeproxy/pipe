package tls

import (
	"context"

	"github.com/pipeproxy/pipe/components/stream"
	"github.com/pipeproxy/pipe/components/tls"
)

type TlsUp struct {
	handler   stream.Handler
	tlsConfig tls.TLS
}

func NewTlsUp(handler stream.Handler, tlsConfig tls.TLS) *TlsUp {
	return &TlsUp{
		handler:   handler,
		tlsConfig: tlsConfig,
	}
}

func (t *TlsUp) ServeStream(ctx context.Context, stm stream.Stream) {
	conn := tls.Client(stm, t.tlsConfig.TLS())
	t.handler.ServeStream(ctx, conn)
}

type TlsDown struct {
	handler   stream.Handler
	tlsConfig tls.TLS
}

func NewTlsDown(handler stream.Handler, tlsConfig tls.TLS) *TlsDown {
	return &TlsDown{
		handler:   handler,
		tlsConfig: tlsConfig,
	}
}

func (t *TlsDown) ServeStream(ctx context.Context, stm stream.Stream) {
	conn := tls.Server(stm, t.tlsConfig.TLS())
	t.handler.ServeStream(ctx, conn)
}
