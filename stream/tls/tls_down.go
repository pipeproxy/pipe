package tls

import (
	"context"
	"crypto/tls"

	"github.com/wzshiming/pipe/stream"
)

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
