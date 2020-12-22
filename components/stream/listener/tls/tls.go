package tls

import (
	"context"

	"github.com/pipeproxy/pipe/components/stream"
	"github.com/pipeproxy/pipe/components/tls"
	"github.com/wzshiming/logger"
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

func (t *Tls) ListenStream(ctx context.Context) (stream.StreamListener, error) {
	log := logger.FromContext(ctx)
	if log.Enabled() {
		log = log.WithName("tls")
		ctx = logger.WithContext(ctx, log)
	}
	listener, err := t.listenConfig.ListenStream(ctx)
	if err != nil {
		return nil, err
	}
	return tls.NewListener(listener, t.tlsConfig.TLS()), nil
}

func (t *Tls) IsVirtual() bool {
	return t.listenConfig.IsVirtual()
}
