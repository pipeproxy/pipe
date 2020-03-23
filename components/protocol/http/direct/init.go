package direct

import (
	"io/ioutil"
	"net/http"

	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/components/protocol/http/template"
	"github.com/wzshiming/pipe/components/stdio/input"
)

const name = "direct"

func init() {
	register.Register(name, NewDirectWithConfig)
}

type Config struct {
	Code int
	Body input.Input
}

func NewDirectWithConfig(conf *Config) (http.Handler, error) {
	body, err := ioutil.ReadAll(conf.Body)
	if err != nil {
		return nil, err
	}
	temp, err := template.NewFormat(string(body))
	if err != nil {
		return nil, err
	}
	return NewDirect(conf.Code, temp), nil
}
