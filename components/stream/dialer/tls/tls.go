package tls

import (
	"context"

	"github.com/pipeproxy/pipe/components/stream"
	"github.com/pipeproxy/pipe/components/tls"
)

type Tls struct {
	dialer    stream.Dialer
	tlsConfig tls.TLS
}

func NewTls(dialer stream.Dialer, tlsConfig tls.TLS) *Tls {
	return &Tls{
		dialer:    dialer,
		tlsConfig: tlsConfig,
	}
}

func (d *Tls) DialStream(ctx context.Context) (stream.Stream, error) {
	stm, err := d.dialer.DialStream(ctx)
	if err != nil {
		return nil, err
	}
	return tls.Client(stm, d.tlsConfig.TLS()), nil
}
