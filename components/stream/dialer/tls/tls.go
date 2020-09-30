package tls

import (
	"context"

	"github.com/wzshiming/pipe/components/stream"
	"github.com/wzshiming/pipe/components/tls"
)

type Tls struct {
	dialer    stream.Dialer
	tlsConfig *tls.Config
}

func NewTls(dialer stream.Dialer, tlsConfig *tls.Config) *Tls {
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
	if d.tlsConfig != nil {
		stm = tls.Client(stm, d.tlsConfig)
	}
	return stm, nil
}
