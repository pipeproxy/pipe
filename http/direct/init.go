package direct

import (
	"net/http"

	"github.com/wzshiming/pipe/configure"
	"github.com/wzshiming/pipe/http/template"
)

const name = "direct"

func init() {
	configure.Register(name, NewDirectWithConfig)
}

type Config struct {
	Code int
	Body string
}

func NewDirectWithConfig(conf *Config) (http.Handler, error) {
	temp, err := template.NewFormat(conf.Body)
	if err != nil {
		return nil, err
	}
	return NewDirect(conf.Code, temp), nil
}
