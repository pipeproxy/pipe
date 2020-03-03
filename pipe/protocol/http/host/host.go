package host

import (
	"fmt"
	"net/http"
	"strings"
)

var (
	ErrNotFound = fmt.Errorf("error not found")
	ErrFormat   = fmt.Errorf("error only a single asterisk is currently supported")
)

// Host is an host multiplexer.
type Host struct {
	domains  map[string]http.Handler
	matchers []*matcher
	notFound http.Handler
}

type matcher struct {
	prefix  string
	suffix  string
	handler http.Handler
}

func NewHost() *Host {
	p := &Host{
		domains: map[string]http.Handler{},
	}
	return p
}

func (h *Host) NotFound(handler http.Handler) error {
	h.notFound = handler
	return nil
}

func (h *Host) Handle(host string, handler http.Handler) error {
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
func (h *Host) Handler(host string) (handler http.Handler, err error) {
	handler, ok := h.domains[host]
	if ok {
		return handler, nil
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
		return m.handler, nil
	}
	if h.notFound == nil {
		return nil, ErrNotFound
	}
	return h.notFound, nil
}

func (h *Host) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	host := r.Host
	handler, err := h.Handler(host)
	if err != nil || handler == nil {
		handler = http.HandlerFunc(http.NotFound)
	}
	handler.ServeHTTP(rw, r)
}
