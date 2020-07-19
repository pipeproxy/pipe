// DO NOT EDIT! Code generated.
package reference

import (
	"net/http"

	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/internal/logger"
)

func init() {
	register.Register("ref", NewHandlerRefWithConfig)
	register.Register("def", NewHandlerDefWithConfig)
	register.Register("none", NewHandlerNone)
}

type Config struct {
	Name string
	Def  http.Handler `json:",omitempty"`
}

func NewHandlerRefWithConfig(conf *Config) (http.Handler, error) {
	o := &Handler{
		Name: conf.Name,
		Def:  conf.Def,
	}
	return o, nil
}

func NewHandlerDefWithConfig(conf *Config) (http.Handler, error) {
	HandlerStore[conf.Name] = conf.Def
	return conf.Def, nil
}

var HandlerStore = map[string]http.Handler{}

func HandlerFind(name string, defaults http.Handler) http.Handler {
	o, ok := HandlerStore[name]
	if ok {
		return o
	}
	if defaults != nil {
		return defaults
	}
	return HandlerNone{}
}

type HandlerNone struct{}

func NewHandlerNone() http.Handler {
	return HandlerNone{}
}

func (HandlerNone) ServeHTTP(_ http.ResponseWriter, _ *http.Request) {
	logger.Warn("this is none of http.Handler")
	return
}

type Handler struct {
	Name string
	Def  http.Handler
}

func (o *Handler) ServeHTTP(responsewriter http.ResponseWriter, b *http.Request) {
	HandlerFind(o.Name, o.Def).ServeHTTP(responsewriter, b)
}
