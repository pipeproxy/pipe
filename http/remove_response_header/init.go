package remove_response_header

import (
	"net/http"

	"github.com/wzshiming/pipe/configure"
)

const name = "remove_response_header"

func init() {
	configure.Register(name, NewRemoveResponseHeaderWithConfig)
}

type Config struct {
	Key string
}

func NewRemoveResponseHeaderWithConfig(conf *Config) http.Handler {
	return NewRemoveResponseHeader(conf.Key)
}
