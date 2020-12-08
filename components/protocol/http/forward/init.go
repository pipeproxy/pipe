package forward

import (
	"net/http"
	"net/url"

	"github.com/pipeproxy/pipe/components/common/register"
	"github.com/pipeproxy/pipe/components/stream"
	"github.com/pipeproxy/pipe/components/stream/dialer/dialer"
)

const (
	name = "forward"
)

func init() {
	register.Register(name, NewForwardWithConfig)
}

type Config struct {
	Dialer stream.Dialer `json:",omitempty"`
	URL    string
}

// NewForwardWithConfig create a new forward with config.
func NewForwardWithConfig(conf *Config) (http.Handler, error) {
	u, err := url.Parse(conf.URL)
	if err != nil {
		return nil, err
	}
	if conf.Dialer == nil {
		conf.Dialer = dialer.NewDialer("tcp", u.Host, false, false)
	}
	return NewForward(conf.URL, conf.Dialer)
}
