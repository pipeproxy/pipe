package tls

import (
	"context"
	"crypto/tls"

	"github.com/wzshiming/pipe/components/stream/dialer"

	"github.com/wzshiming/pipe/components/stream"
)

type Tls struct {
	dialer    dialer.Dialer
	tlsConfig *tls.Config
}

func NewTls(dialer dialer.Dialer, tlsConfig *tls.Config) *Tls {
	return &Tls{
		dialer:    dialer,
		tlsConfig: tlsConfig,
	}
}

func (t *Tls) DialStream(ctx context.Context) (stream.Stream, error) {
	stm, err := t.dialer.DialStream(ctx)
	if err != nil {
		return nil, err
	}
	stm = tls.Client(stm, t.tlsConfig)
	return stm, nil
}
