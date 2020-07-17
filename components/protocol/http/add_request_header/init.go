package add_request_header

import (
	"net/http"

	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/internal/http/template"
)

const (
	name = "add_request_header"
)

func init() {
	register.Register(name, NewAddRequestHeaderWithConfig)
}

type Config struct {
	Key   string
	Value string
}

func NewAddRequestHeaderWithConfig(conf *Config) (http.Handler, error) {
	temp, err := template.NewFormat(conf.Value)
	if err != nil {
		return nil, err
	}
	return NewAddRequestHeader(conf.Key, temp), nil
}
