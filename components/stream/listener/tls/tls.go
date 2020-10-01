package tls

import (
	"context"

	"github.com/wzshiming/pipe/components/stream"
	"github.com/wzshiming/pipe/components/tls"
)

type Tls struct {
	listenConfig stream.ListenConfig
	tlsConfig    tls.TLS
}

func NewTls(listenConfig stream.ListenConfig, tlsConfig tls.TLS) *Tls {
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
	return tls.NewListener(listener, d.tlsConfig.TLS()), nil
}
