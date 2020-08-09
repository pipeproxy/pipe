package strip_prefix

import (
	"net/http"

	"github.com/wzshiming/pipe/components/common/register"
)

const (
	name = "strip_prefix"
)

func init() {
	register.Register(name, NewStripPrefixWithConfig)
}

type Config struct {
	Prefix  string
	Handler http.Handler
}

func NewStripPrefixWithConfig(conf *Config) http.Handler {
	return http.StripPrefix(conf.Prefix, conf.Handler)
}
