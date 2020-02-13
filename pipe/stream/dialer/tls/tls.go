package tls

import (
	"context"
	"crypto/tls"

	"github.com/wzshiming/pipe/pipe/stream/dialer"

	"github.com/wzshiming/pipe/pipe/stream"
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

func (t *Tls) Dial(ctx context.Context) (stream.Stream, error) {
	stm, err := t.dialer.Dial(ctx)
	if err != nil {
		return nil, err
	}
	stm = tls.Client(stm, t.tlsConfig)
	return stm, nil
}
