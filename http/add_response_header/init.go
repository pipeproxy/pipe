package add_response_header

import (
	"net/http"

	"github.com/wzshiming/pipe/configure"
)

const name = "add_response_header"

func init() {
	configure.Register(name, NewAddResponseHeaderWithConfig)
}

type Config struct {
	Key    string
	Values []string
}

func NewAddResponseHeaderWithConfig(conf *Config) http.Handler {
	return NewAddResponseHeader(conf.Key, conf.Values)
}
