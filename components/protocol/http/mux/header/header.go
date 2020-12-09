package header

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/pipeproxy/pipe/internal/http/template"
)

// Header is an host multiplexer.
type Header struct {
	routes   []route
	notFound http.Handler
}

type route struct {
	matches []Matcher
	handler http.Handler
}

type Matcher interface {
	Match(v http.Header) bool
}

// NewHeader create a new Header.
func NewHeader() *Header {
	p := &Header{}
	return p
}

// NotFound replies to the handler with an Handler not found error.
func (h *Header) NotFound(handler http.Handler) {
	h.notFound = handler
}

func (h *Header) Handle(matches []Matcher, handler http.Handler) {
	h.routes = append(h.routes, route{
		matches: matches,
		handler: handler,
	})
}

func (h *Header) Handler(v http.Header) http.Handler {
	for _, route := range h.routes {
		match := true
		for _, matcher := range route.matches {
			if !matcher.Match(v) {
				match = false
				break
			}
		}
		if match {
			return route.handler
		}
	}
	if h.notFound == nil {
		return template.NotFoundHandler
	}
	return h.notFound
}

func (h *Header) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.Handler(r.Header).ServeHTTP(rw, r)
}

type Exact struct {
	Key   string
	Value string
}

func (e *Exact) Match(v http.Header) bool {
	return v.Get(e.Key) == e.Value
}

type Present struct {
	Key string
}

func (p *Present) Match(v http.Header) bool {
	_, ok := v[p.Key]
	return ok
}

type Regexp struct {
	Key   string
	Value *regexp.Regexp
}

func (r *Regexp) Match(v http.Header) bool {
	return r.Value.MatchString(v.Get(r.Key))
}

type Prefix struct {
	Key   string
	Value string
}

func (p *Prefix) Match(v http.Header) bool {
	return strings.HasPrefix(v.Get(p.Key), p.Value)
}

type Suffix struct {
	Key   string
	Value string
}

func (s *Suffix) Match(v http.Header) bool {
	return strings.HasSuffix(v.Get(s.Key), s.Value)
}

type Contains struct {
	Key   string
	Value string
}

func (s *Contains) Match(v http.Header) bool {
	return strings.Contains(v.Get(s.Key), s.Value)
}
