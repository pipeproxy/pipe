package add_request_header

import (
	"net/http"

	"github.com/wzshiming/pipe/configure/manager"
	"github.com/wzshiming/pipe/pipe/protocol/http/template"
)

const name = "add_request_header"

func init() {
	manager.Register(name, NewAddRequestHeaderWithConfig)
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
