package remove_request_header

import (
	"net/http"

	"github.com/wzshiming/pipe/components/common/register"
)

const (
	name = "remove_request_header"
)

func init() {
	register.Register(name, NewRemoveRequestHeaderWithConfig)
}

type Config struct {
	Key string
}

func NewRemoveRequestHeaderWithConfig(conf *Config) http.Handler {
	return NewRemoveRequestHeader(conf.Key)
}
