package tls

import (
	"context"

	"github.com/wzshiming/pipe/components/stream"
	"github.com/wzshiming/pipe/components/tls"
)

type TlsUp struct {
	handler   stream.Handler
	tlsConfig *tls.Config
}

func NewTlsUp(handler stream.Handler, tlsConfig *tls.Config) *TlsUp {
	return &TlsUp{
		handler:   handler,
		tlsConfig: tlsConfig,
	}
}

func (t *TlsUp) ServeStream(ctx context.Context, stm stream.Stream) {
	conn := tls.Client(stm, t.tlsConfig)
	t.handler.ServeStream(ctx, conn)
}

type TlsDown struct {
	handler   stream.Handler
	tlsConfig *tls.Config
}

func NewTlsDown(handler stream.Handler, tlsConfig *tls.Config) *TlsDown {
	return &TlsDown{
		handler:   handler,
		tlsConfig: tlsConfig,
	}
}

func (t *TlsDown) ServeStream(ctx context.Context, stm stream.Stream) {
	conn := tls.Server(stm, t.tlsConfig)
	t.handler.ServeStream(ctx, conn)
}
