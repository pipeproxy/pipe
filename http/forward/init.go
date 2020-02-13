package forward

import (
	"context"
	"net"
	"net/http"

	"github.com/wzshiming/pipe/configure/manager"
	"github.com/wzshiming/pipe/dialer"
	"github.com/wzshiming/pipe/tls"
)

const name = "forward"

func init() {
	manager.Register(name, NewForwardWithConfig)
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
			return conf.Dialer.Dial(ctx)
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
