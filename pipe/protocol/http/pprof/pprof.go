package pprof

import (
	"net/http"
	"net/http/pprof"
	"strings"
)

type Pprof struct {
}

func NewPprof() Pprof {
	return Pprof{}
}

func (Pprof) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	index := strings.LastIndexByte(path, '/')
	if index != -1 {
		path = path[index+1:]
	}
	switch path {
	case "", "index":
		pprof.Index(rw, r)
	case "cmdline":
		pprof.Cmdline(rw, r)
	case "profile":
		pprof.Profile(rw, r)
	case "symbol":
		pprof.Symbol(rw, r)
	case "trace":
		pprof.Trace(rw, r)
	default:
		pprof.Handler(path).ServeHTTP(rw, r)
	}
}
