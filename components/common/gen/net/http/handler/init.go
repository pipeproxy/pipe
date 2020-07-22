// DO NOT EDIT! Code generated.
package reference

import (
	"net/http"
	"sync"

	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/internal/logger"
)

func init() {
	register.Register("ref", NewHandlerRefWithConfig)
	register.Register("def", NewHandlerDefWithConfig)
	register.Register("none", newHandlerNone)
}

type Config struct {
	Name string
	Def  http.Handler `json:",omitempty"`
}

func NewHandlerRefWithConfig(conf *Config) http.Handler {
	o := &Handler{
		Name: conf.Name,
		Def:  conf.Def,
	}
	return o
}

func NewHandlerDefWithConfig(conf *Config) http.Handler {
	return HandlerPut(conf.Name, conf.Def)
}

var (
	mut           sync.RWMutex
	_HandlerStore = map[string]http.Handler{}
)

func HandlerPut(name string, def http.Handler) http.Handler {
	if def == nil {
		def = HandlerNone
	}
	mut.Lock()
	_HandlerStore[name] = def
	mut.Unlock()
	return def
}

func HandlerGet(name string, defaults http.Handler) http.Handler {
	mut.RLock()
	o, ok := _HandlerStore[name]
	mut.RUnlock()
	if ok {
		return o
	}
	if defaults != nil {
		return defaults
	}
	return HandlerNone
}

var HandlerNone _HandlerNone

type _HandlerNone struct{}

func newHandlerNone() http.Handler {
	return HandlerNone
}

func (_HandlerNone) ServeHTTP(_ http.ResponseWriter, _ *http.Request) {
	logger.Warn("this is none of http.Handler")
	return
}

type Handler struct {
	Name string
	Def  http.Handler
}

func (o *Handler) ServeHTTP(responsewriter http.ResponseWriter, b *http.Request) {
	HandlerGet(o.Name, o.Def).ServeHTTP(responsewriter, b)
}
