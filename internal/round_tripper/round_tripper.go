package round_tripper

import (
	"context"
	"net/http"

	"github.com/pipeproxy/pipe/components/stream"
)

var defaultTransport = http.DefaultTransport.(*http.Transport).Clone()

func RoundTripperList(ds []stream.Dialer) []http.RoundTripper {
	out := make([]http.RoundTripper, 0, len(ds))
	for _, d := range ds {
		out = append(out, RoundTripper(d))
	}
	return out
}

func RoundTripper(d stream.Dialer) http.RoundTripper {
	if d == nil {
		return defaultTransport
	}
	if r, ok := d.(isRoundTripper); ok {
		return r.RoundTripper()
	}
	transport := defaultTransport.Clone()
	if c, ok := d.(isTCP); ok && c.IsTCP() {
		transport.DialTLSContext = func(ctx context.Context, network, addr string) (stream.Stream, error) {
			return d.DialStream(ctx)
		}
	} else {
		transport.DialContext = func(ctx context.Context, network, addr string) (stream.Stream, error) {
			return d.DialStream(ctx)
		}
	}
	return transport
}

type isRoundTripper interface {
	RoundTripper() http.RoundTripper
}

type isTCP interface {
	IsTCP() bool
}
