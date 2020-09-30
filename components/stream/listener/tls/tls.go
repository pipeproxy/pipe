package tls

import (
	"context"

	"github.com/wzshiming/pipe/components/stream"
	"github.com/wzshiming/pipe/components/tls"
)

type Tls struct {
	listenConfig stream.ListenConfig
	tlsConfig    *tls.Config
}

func NewTls(listenConfig stream.ListenConfig, tlsConfig *tls.Config) *Tls {
	return &Tls{
		listenConfig: listenConfig,
		tlsConfig:    tlsConfig,
	}
}

func (d *Tls) ListenStream(ctx context.Context) (stream.StreamListener, error) {
	listener, err := d.listenConfig.ListenStream(ctx)
	if err != nil {
		return nil, err
	}
	return tls.NewListener(listener, d.tlsConfig), nil
}
