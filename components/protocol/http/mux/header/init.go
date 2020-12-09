package header

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/pipeproxy/pipe/components/common/register"
)

const (
	name = "header"
)

func init() {
	register.Register(name, NewHeaderWithConfig)
}

var (
	ErrNotHandler = fmt.Errorf("error not handler")
	ErrNotRouter  = fmt.Errorf("error not router")
)

type RouteMatch struct {
	Key      string
	Exact    string `json:",omitempty"`
	Present  bool   `json:",omitempty"`
	Regexp   string `json:",omitempty"`
	Prefix   string `json:",omitempty"`
	Suffix   string `json:",omitempty"`
	Contains string `json:",omitempty"`
}

type Route struct {
	Matches []RouteMatch
	Handler http.Handler
}

type Config struct {
	Headers  []*Route
	NotFound http.Handler `json:",omitempty"`
}

func NewHeaderWithConfig(conf *Config) (http.Handler, error) {
	mux := NewHeader()
	mux.NotFound(conf.NotFound)
	for _, route := range conf.Headers {
		if route.Handler == nil {
			return nil, ErrNotHandler
		}
		if len(route.Matches) == 0 {
			return nil, ErrNotRouter
		}
		var matchers []Matcher

		for _, match := range route.Matches {
			if match.Exact != "" {
				matchers = append(matchers, &Exact{
					Key:   match.Key,
					Value: match.Exact,
				})
			}
			if match.Present {
				matchers = append(matchers, &Present{
					Key: match.Key,
				})
			}
			if match.Regexp != "" {
				r, err := regexp.Compile(match.Regexp)
				if err != nil {
					return nil, err
				}
				matchers = append(matchers, &Regexp{
					Key:   match.Key,
					Value: r,
				})
			}
			if match.Prefix != "" {
				matchers = append(matchers, &Prefix{
					Key:   match.Key,
					Value: match.Prefix,
				})
			}
			if match.Suffix != "" {
				matchers = append(matchers, &Suffix{
					Key:   match.Key,
					Value: match.Suffix,
				})
			}
			if match.Contains != "" {
				matchers = append(matchers, &Contains{
					Key:   match.Key,
					Value: match.Contains,
				})
			}
		}
		mux.Handle(matchers, route.Handler)
	}
	return mux, nil
}
