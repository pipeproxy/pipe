package add_request_header

import (
	"net/http"

	"github.com/wzshiming/pipe/configure"
)

const name = "add_request_header"

func init() {
	configure.Register(name, NewAddRequestHeaderWithConfig)
}

type Config struct {
	Key    string
	Values []string
}

func NewAddRequestHeaderWithConfig(conf *Config) http.Handler {
	return NewAddRequestHeader(conf.Key, conf.Values)
}
