package tls

import (
	"context"
	"crypto/tls"

	"github.com/wzshiming/pipe/components/stream"
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
