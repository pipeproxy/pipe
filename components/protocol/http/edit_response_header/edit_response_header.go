package edit_response_header

import (
	"net/http"

	"github.com/pipeproxy/pipe/internal/http/template"
)

type pair struct {
	key   string
	value template.Format
}

type EditResponseHeader struct {
	del []string
	set []pair
	add []pair
}

func NewEditResponseHeader(del []string, set []Pair, add []Pair) (*EditResponseHeader, error) {
	e := &EditResponseHeader{
		del: del,
		set: make([]pair, 0, len(set)),
		add: make([]pair, 0, len(add)),
	}
	for _, s := range set {
		temp, err := template.NewFormat(s.Value)
		if err != nil {
			return nil, err
		}
		e.set = append(e.set, pair{
			key:   s.Key,
			value: temp,
		})
	}
	for _, a := range add {
		temp, err := template.NewFormat(a.Value)
		if err != nil {
			return nil, err
		}
		e.add = append(e.add, pair{
			key:   a.Key,
			value: temp,
		})
	}
	return e, nil
}

func (e *EditResponseHeader) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	header := rw.Header()
	for _, k := range e.del {
		header.Del(k)
	}
	for _, s := range e.set {
		header.Set(s.key, s.value.FormatString(r))
	}
	for _, a := range e.add {
		header.Add(a.key, a.value.FormatString(r))
	}
}
