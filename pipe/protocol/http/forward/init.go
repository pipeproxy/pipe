package forward

import (
	"context"
	"net"
	"net/http"

	"github.com/wzshiming/pipe/configure/decode"
	"github.com/wzshiming/pipe/pipe/stream/dialer"
	"github.com/wzshiming/pipe/pipe/tls"
)

const name = "forward"

func init() {
	decode.Register(name, NewForwardWithConfig)
}

type Config struct {
	TLS    tls.TLS
	Dialer dialer.Dialer
	Pass   string
}

var defaultTransport = http.DefaultTransport.(*http.Transport)

// NewForwardWithConfig create a new forward with config.
func NewForwardWithConfig(conf *Config) (http.Handler, error) {
	transport := defaultTransport
	if conf.Dialer != nil {
		transport = defaultTransport.Clone()
		transport.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
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

	return NewForward(conf.Pass, transport)
}
