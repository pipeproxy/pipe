package remove_response_header

import (
	"net/http"

	"github.com/wzshiming/pipe/components/common/register"
)

const (
	name = "remove_response_header"
)

func init() {
	register.Register(name, NewRemoveResponseHeaderWithConfig)
}

type Config struct {
	Key string
}

func NewRemoveResponseHeaderWithConfig(conf *Config) http.Handler {
	return NewRemoveResponseHeader(conf.Key)
}
