package direct

import (
	"io/ioutil"
	"net/http"

	"github.com/wzshiming/pipe/configure/decode"
	"github.com/wzshiming/pipe/pipe/protocol/http/template"
	"github.com/wzshiming/pipe/pipe/stdio/input"
)

const name = "direct"

func init() {
	decode.Register(name, NewDirectWithConfig)
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
	err = conf.Body.Close()
	if err != nil {
		return nil, err
	}

	temp, err := template.NewFormat(string(body))
	if err != nil {
		return nil, err
	}
	return NewDirect(conf.Code, temp), nil
}
