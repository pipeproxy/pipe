package query

import (
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/pipeproxy/pipe/internal/http/template"
)

// Query is an host multiplexer.
type Query struct {
	routes   []route
	notFound http.Handler
}

type route struct {
	matches []Matcher
	handler http.Handler
}

type Matcher interface {
	Match(v url.Values) bool
}

// NewQuery create a new Query.
func NewQuery() *Query {
	p := &Query{}
	return p
}

// NotFound replies to the handler with an Handler not found error.
func (q *Query) NotFound(handler http.Handler) {
	q.notFound = handler
}

func (q *Query) Handle(matches []Matcher, handler http.Handler) {
	q.routes = append(q.routes, route{
		matches: matches,
		handler: handler,
	})
}

func (q *Query) Handler(v url.Values) http.Handler {
	for _, route := range q.routes {
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
	if q.notFound == nil {
		return template.NotFoundHandler
	}
	return q.notFound
}

func (q *Query) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	q.Handler(r.URL.Query()).ServeHTTP(rw, r)
}

type Exact struct {
	Key   string
	Value string
}

func (e *Exact) Match(v url.Values) bool {
	return v.Get(e.Key) == e.Value
}

type Present struct {
	Key string
}

func (p *Present) Match(v url.Values) bool {
	_, ok := v[p.Key]
	return ok
}

type Regexp struct {
	Key   string
	Value *regexp.Regexp
}

func (r *Regexp) Match(v url.Values) bool {
	return r.Value.MatchString(v.Get(r.Key))
}

type Prefix struct {
	Key   string
	Value string
}

func (p *Prefix) Match(v url.Values) bool {
	return strings.HasPrefix(v.Get(p.Key), p.Value)
}

type Suffix struct {
	Key   string
	Value string
}

func (s *Suffix) Match(v url.Values) bool {
	return strings.HasSuffix(v.Get(s.Key), s.Value)
}

type Contains struct {
	Key   string
	Value string
}

func (s *Contains) Match(v url.Values) bool {
	return strings.Contains(v.Get(s.Key), s.Value)
}
