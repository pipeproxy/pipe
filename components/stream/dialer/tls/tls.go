package tls

import (
	"context"
	"strings"

	"github.com/pipeproxy/pipe/components/balance"
	"github.com/pipeproxy/pipe/components/stream"
	"github.com/pipeproxy/pipe/components/tls"
	"github.com/wzshiming/logger"
)

type Tls struct {
	dialer    stream.Dialer
	tlsConfig tls.TLS
	name      string
}

func NewTls(dialer stream.Dialer, tlsConfig tls.TLS) *Tls {
	t := &Tls{
		dialer:    dialer,
		tlsConfig: tlsConfig,
	}
	t.name = t.getName()
	return t
}

func (t *Tls) DialStream(ctx context.Context) (stream.Stream, error) {
	log := logger.FromContext(ctx)
	log = log.WithName("tls")
	ctx = logger.WithContext(ctx, log)
	stm, err := t.dialer.DialStream(ctx)
	if err != nil {
		return nil, err
	}
	return tls.Client(stm, t.tlsConfig.TLS()), nil
}

func (t *Tls) Targets() []stream.Dialer {
	ts := t.dialer.Targets()
	ds := make([]stream.Dialer, 0, len(ts))
	for _, target := range ts {
		ds = append(ds, NewTls(target, t.tlsConfig))
	}
	return ds
}

func (t *Tls) Policy() balance.Policy {
	return t.dialer.Policy()
}

func (t *Tls) String() string {
	return t.name
}

func (t *Tls) getName() string {
	return strings.Join([]string{"tls", t.dialer.String()}, "://")
}
