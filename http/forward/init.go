package forward

import (
	"context"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/wzshiming/pipe/configure/manager"
	"github.com/wzshiming/pipe/dialer"
	"github.com/wzshiming/pipe/internal/pool"
)

const name = "forward"

func init() {
	manager.Register(name, NewForwardWithConfig)
}

type Config struct {
	Dialer dialer.Dialer
	Pass   string
}

var defaultTransport = http.Transport{
	DialContext: (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
		DualStack: true,
	}).DialContext,
	ForceAttemptHTTP2:     true,
	MaxIdleConns:          100,
	IdleConnTimeout:       90 * time.Second,
	TLSHandshakeTimeout:   10 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
}

// NewForwardWithConfig create a new forward with config.
func NewForwardWithConfig(conf *Config) (http.Handler, error) {
	u, err := url.Parse(conf.Pass)
	if err != nil {
		return nil, err
	}

	rp := httputil.NewSingleHostReverseProxy(u)
	rp.Transport = &defaultTransport
	if conf.Dialer != nil {
		rp.BufferPool = pool.Buffer
		rp.Transport = &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				return conf.Dialer.Dial(ctx)
			},
			ForceAttemptHTTP2:     defaultTransport.ForceAttemptHTTP2,
			MaxIdleConns:          defaultTransport.MaxIdleConns,
			IdleConnTimeout:       defaultTransport.IdleConnTimeout,
			TLSHandshakeTimeout:   defaultTransport.TLSHandshakeTimeout,
			ExpectContinueTimeout: defaultTransport.ExpectContinueTimeout,
		}
	}
	return rp, nil
}
