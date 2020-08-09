package forward

import (
	"net/http"

	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/components/protocol/http/round_tripper"
)

const (
	name = "forward"
)

func init() {
	register.Register(name, NewForwardWithConfig)
}

type Config struct {
	RoundTripper round_tripper.RoundTripper `json:",omitempty"`
	URL          string
}

// NewForwardWithConfig create a new forward with config.
func NewForwardWithConfig(conf *Config) (http.Handler, error) {
	roundTripper := conf.RoundTripper
	if roundTripper == nil {
		roundTripper = http.DefaultTransport
	}
	return NewForward(conf.URL, roundTripper)
}
