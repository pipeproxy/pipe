package round_tripper

import (
	"context"
	"net/http"

	"github.com/pipeproxy/pipe/components/stream"
)

var defaultTransport = http.DefaultTransport.(*http.Transport).Clone()

func roundTrippers(ds []stream.Dialer) []http.RoundTripper {
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
	ds := d.Targets()
	if len(ds) > 1 {
		return NewLB(d.Policy(), roundTrippers(ds))
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

type isTCP interface {
	IsTCP() bool
}
