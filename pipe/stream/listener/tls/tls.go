package tls

import (
	"context"
	"crypto/tls"
	"net"

	"github.com/wzshiming/pipe/pipe/stream/listener"
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

func (t *Tls) Listen(ctx context.Context) (net.Listener, error) {
	listener, err := t.listenConfig.Listen(ctx)
	if err != nil {
		return nil, err
	}
	return tls.NewListener(listener, t.tlsConfig), nil
}
