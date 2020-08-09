package transport

import (
	"context"
	"net/http"

	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/components/protocol/http/round_tripper"
	"github.com/wzshiming/pipe/components/stream"
	"github.com/wzshiming/pipe/components/tls"
)

const (
	name = "transport"
)

func init() {
	register.Register(name, NewTransportWithConfig)
}

type Config struct {
	TLS    tls.TLS
	Dialer stream.Dialer
}

var defaultTransport = http.DefaultTransport.(*http.Transport)

func NewTransportWithConfig(conf *Config) (round_tripper.RoundTripper, error) {
	transport := defaultTransport
	if conf.Dialer != nil {
		transport = defaultTransport.Clone()
		transport.DialContext = func(ctx context.Context, network, addr string) (stream.Stream, error) {
			return conf.Dialer.DialStream(ctx)
		}
		if conf.TLS != nil {
			transport.TLSClientConfig = conf.TLS.TLS()
		}
	} else {
		if conf.TLS != nil {
			transport = defaultTransport.Clone()
			transport.TLSClientConfig = conf.TLS.TLS()
		}
	}
	return transport, nil
}
