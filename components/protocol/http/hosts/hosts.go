package hosts

import (
	"fmt"
	"net/http"
	"strings"
)

var (
	ErrFormat = fmt.Errorf("error only a single asterisk is currently supported")
)

// Hosts is an hosts multiplexer.
type Hosts struct {
	domains  map[string]http.Handler
	matchers []*matcher
	notFound http.Handler
}

type matcher struct {
	prefix  string
	suffix  string
	handler http.Handler
}

func NewHosts() *Hosts {
	p := &Hosts{
		domains: map[string]http.Handler{},
	}
	return p
}

func (h *Hosts) NotFound(handler http.Handler) {
	h.notFound = handler
}

func (h *Hosts) Handle(host string, handler http.Handler) error {
	if host == "*" {
		h.NotFound(handler)
		return nil
	}
	split := strings.Split(host, "*")
	switch len(split) {
	default:
		return ErrFormat
	case 1:
		h.domains[split[0]] = handler
	case 2:
		m := &matcher{
			prefix: split[0],
			suffix: split[1],
		}
		h.matchers = append(h.matchers, m)
	}
	return nil
}

// handler returns most matching handler and prefix bytes data to use for the given reader.
func (h *Hosts) Handler(host string) (handler http.Handler) {
	handler, ok := h.domains[host]
	if ok {
		return handler
	}
	for _, m := range h.matchers {
		if m.prefix != "" {
			if !strings.HasPrefix(host, m.prefix) {
				continue
			}
		}
		if m.suffix != "" {
			if !strings.HasSuffix(host, m.suffix) {
				continue
			}
		}
		return m.handler
	}
	if h.notFound == nil {
		return http.HandlerFunc(http.NotFound)
	}
	return h.notFound
}

func (h *Hosts) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.Handler(r.Host).ServeHTTP(rw, r)
}
