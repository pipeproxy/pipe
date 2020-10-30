package round_tripper

import (
	"context"
	"net/http"

	"github.com/pipeproxy/pipe/components/stream"
)

var defaultTransport = http.DefaultTransport.(*http.Transport).Clone()

func RoundTripper(d stream.Dialer) http.RoundTripper {
	if d == nil {
		return defaultTransport
	}
	transport := defaultTransport.Clone()
	if c, ok := d.(isTCP); ok && c.IsTCP() {
		transport.DialContext = func(ctx context.Context, network, addr string) (stream.Stream, error) {
			return d.DialStream(ctx)
		}
	} else {
		transport.DialContext = func(ctx context.Context, network, addr string) (stream.Stream, error) {
			return d.DialStream(ctx)
		}
	}
	return transport
}

type isTCP interface {
	IsTCP() bool
}
