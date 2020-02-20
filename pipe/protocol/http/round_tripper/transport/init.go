package transport

import (
	"context"
	"net/http"

	"github.com/wzshiming/pipe/configure/decode"
	"github.com/wzshiming/pipe/pipe/protocol/http/round_tripper"
	"github.com/wzshiming/pipe/pipe/stream"
	"github.com/wzshiming/pipe/pipe/stream/dialer"
	"github.com/wzshiming/pipe/pipe/tls"
)

const name = "transport"

func init() {
	decode.Register(name, NewTransportWithConfig)
}

type Config struct {
	TLS    tls.TLS
	Dialer dialer.Dialer
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
