package direct

import (
	"io/ioutil"
	"net/http"

	"github.com/pipeproxy/pipe/components/common/register"
	"github.com/pipeproxy/pipe/components/stdio/input"
	"github.com/pipeproxy/pipe/internal/http/template"
)

const (
	name = "direct"
)

func init() {
	register.Register(name, NewDirectWithConfig)
}

type Config struct {
	Code int
	Body input.Input
}

func NewDirectWithConfig(conf *Config) (http.Handler, error) {
	var temp template.Format
	if conf.Body != nil {
		body, err := ioutil.ReadAll(conf.Body)
		if err != nil {
			return nil, err
		}
		temp, err = template.NewFormat(string(body))
		if err != nil {
			return nil, err
		}
	}
	return NewDirect(conf.Code, temp), nil
}
