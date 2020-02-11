package redirect

import (
	"net/http"

	"github.com/wzshiming/pipe/configure/manager"
	"github.com/wzshiming/pipe/http/template"
)

const name = "redirect"

func init() {
	manager.Register(name, NewRedirectWithConfig)
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
