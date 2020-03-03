package mux

import (
	"encoding/binary"
	"fmt"
	"net/http"
	"regexp"
	"sync/atomic"

	"github.com/wzshiming/pipe/internal/pool"
	"github.com/wzshiming/trie"
)

var (
	ErrNotFound           = fmt.Errorf("error not found")
	ErrRouteAlreadyExists = fmt.Errorf("error route already exists")
)

const handlerMapPrefixSize = 4

// Mux is an Applicative protocol multiplexer.
type Mux struct {
	trie         *trie.Trie
	prefixLength int
	size         uint32
	handlers     map[uint32]*regexpRoutes
	paths        map[string]http.Handler
	notFound     http.Handler
}

type regexpRoutes struct {
	Regexps []*regexpRoute
	Handler http.Handler
}

type regexpRoute struct {
	Regexp  *regexp.Regexp
	Handler http.Handler
}

// NewMux create a new Mux.
func NewMux() *Mux {
	p := &Mux{
		trie:     trie.NewTrie(),
		handlers: map[uint32]*regexpRoutes{},
		paths:    map[string]http.Handler{},
	}
	return p
}

// NotFound replies to the handler with an Handler not found error.
func (m *Mux) NotFound(handler http.Handler) error {
	m.notFound = handler
	return nil
}

func (m *Mux) HandlePrefix(prefix string, handler http.Handler) error {
	buf, err := m.setHandler(handler, nil)
	if err != nil {
		return err
	}
	m.handle(prefix, buf)
	return nil
}

func (m *Mux) HandlePrefixAndRegexp(prefix, reg string, handler http.Handler) error {
	r, err := regexp.Compile(reg)
	if err != nil {
		return err
	}
	buf, err := m.setHandler(handler, r)
	if err != nil {
		return err
	}
	m.handle(prefix, buf)
	return nil
}

func (m *Mux) HandlePath(path string, handler http.Handler) error {
	m.paths[path] = handler
	return nil
}

// Handler returns most matching handler and prefix bytes data to use for the given reader.
func (m *Mux) Handler(path string) (handler http.Handler, err error) {
	handler, ok := m.paths[path]
	if ok {
		return handler, nil
	}
	if m.prefixLength == 0 {
		return nil, ErrNotFound
	}

	buf := pool.Buffer.Get()
	defer pool.Buffer.Put(buf)
	i := copy(buf, path[:])
	data, _, _ := m.trie.Mapping().Get(buf[:i])
	if len(data) != 0 {
		conn, ok := m.getHandler(data)
		if ok {
			handler = conn
		}
	}
	if handler != nil {
		return handler, nil
	}
	if m.notFound == nil {
		return nil, ErrNotFound
	}
	return m.notFound, nil
}

func (m *Mux) handle(prefix string, buf []byte) {
	m.trie.Put([]byte(prefix), buf)
	if m.prefixLength < len(prefix) {
		m.prefixLength = len(prefix)
	}
}

func (m *Mux) setHandler(hand http.Handler, reg *regexp.Regexp) ([]byte, error) {
	k := atomic.AddUint32(&m.size, 1)
	buf := make([]byte, handlerMapPrefixSize)
	binary.BigEndian.PutUint32(buf, k)

	_, ok := m.handlers[k]
	if !ok {
		m.handlers[k] = &regexpRoutes{}
	}
	if reg == nil {
		if m.handlers[k].Handler != nil {
			return nil, ErrRouteAlreadyExists
		}
		m.handlers[k].Handler = hand
	} else {
		m.handlers[k].Regexps = append(m.handlers[k].Regexps, &regexpRoute{
			Regexp:  reg,
			Handler: hand,
		})
	}

	return buf, nil
}

func (m *Mux) getHandler(index []byte) (http.Handler, bool) {
	c, ok := m.handlers[binary.BigEndian.Uint32(index)]
	if !ok {
		return nil, false
	}
	if len(c.Regexps) != 0 {
		for _, r := range c.Regexps {
			if r.Regexp.Match(index[handlerMapPrefixSize:]) {
				return r.Handler, true
			}
		}
	}
	if c.Handler == nil {
		return nil, false
	}
	return c.Handler, true
}

func (m *Mux) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	handler, err := m.Handler(path)
	if err != nil || handler == nil {
		handler = http.HandlerFunc(http.NotFound)
	}
	handler.ServeHTTP(rw, r)
}
