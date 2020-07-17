package add_response_header

import (
	"net/http"

	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/internal/http/template"
)

const (
	name = "add_response_header"
)

func init() {
	register.Register(name, NewAddResponseHeaderWithConfig)
}

type Config struct {
	Key   string
	Value string
}

func NewAddResponseHeaderWithConfig(conf *Config) (http.Handler, error) {
	temp, err := template.NewFormat(conf.Value)
	if err != nil {
		return nil, err
	}
	return NewAddResponseHeader(conf.Key, temp), nil
}
