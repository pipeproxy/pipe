package tls

import (
	"context"
	"crypto/tls"

	"github.com/wzshiming/pipe/components/stream/listener"
)

type Tls struct {
	listenConfig listener.ListenConfig
	tlsConfig    *tls.Config
}

func NewTls(listenConfig listener.ListenConfig, tlsConfig *tls.Config) *Tls {
	return &Tls{
		listenConfig: listenConfig,
		tlsConfig:    tlsConfig,
	}
}

func (t *Tls) ListenStream(ctx context.Context) (listener.StreamListener, error) {
	listener, err := t.listenConfig.ListenStream(ctx)
	if err != nil {
		return nil, err
	}
	return tls.NewListener(listener, t.tlsConfig), nil
}
