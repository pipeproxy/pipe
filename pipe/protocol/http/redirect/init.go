package redirect

import (
	"net/http"

	"github.com/wzshiming/pipe/configure/decode"
	"github.com/wzshiming/pipe/pipe/protocol/http/template"
)

const name = "redirect"

func init() {
	decode.Register(name, NewRedirectWithConfig)
}

type Config struct {
	Code     int
	Location string
}

func NewRedirectWithConfig(conf *Config) (http.Handler, error) {
	temp, err := template.NewFormat(conf.Location)
	if err != nil {
		return nil, err
	}
	code := conf.Code
	if code == 0 {
		code = http.StatusFound
	}
	return NewRedirect(code, temp), nil
}
