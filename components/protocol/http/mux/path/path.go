package path

import (
	"encoding/binary"
	"fmt"
	"net/http"
	"regexp"
	"sync/atomic"

	"github.com/pipeproxy/pipe/internal/pool"
	"github.com/wzshiming/trie"
)

var (
	ErrRouteAlreadyExists = fmt.Errorf("error route already exists")
)

const handlerMapPrefixSize = 4

// Path is an host multiplexer.
type Path struct {
	trie         *trie.Trie
	prefixLength int
	size         uint32
	handlers     map[uint32]*regexpRoutes
	paths        map[string]http.Handler
	notFound     http.Handler
}

type regexpRoutes struct {
	matchers []*matcher
	handler  http.Handler
}

type matcher struct {
	match   *regexp.Regexp
	handler http.Handler
}

// NewPath create a new Path.
func NewPath() *Path {
	p := &Path{
		trie:     trie.NewTrie(),
		handlers: map[uint32]*regexpRoutes{},
		paths:    map[string]http.Handler{},
	}
	return p
}

// NotFound replies to the handler with an Handler not found error.
func (m *Path) NotFound(handler http.Handler) {
	m.notFound = handler
	return
}

func (m *Path) HandlePrefix(prefix string, handler http.Handler) error {
	buf, err := m.setHandler(handler, nil)
	if err != nil {
		return err
	}
	m.handle(prefix, buf)
	return nil
}

func (m *Path) HandlePrefixAndRegexp(prefix, reg string, handler http.Handler) error {
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

func (m *Path) HandlePath(path string, handler http.Handler) {
	m.paths[path] = handler
	return
}

// Handler returns most matching handler and prefix bytes data to use for the given reader.
func (m *Path) Handler(path string) (handler http.Handler) {
	handler, ok := m.paths[path]
	if ok {
		return handler
	}
	if m.prefixLength == 0 {
		return http.HandlerFunc(http.NotFound)
	}

	buf := pool.GetBytes()
	defer pool.PutBytes(buf)
	i := copy(buf, path[:])
	data, _, _ := m.trie.Mapping().Get(buf[:i])
	if len(data) != 0 {
		conn, ok := m.getHandler(data)
		if ok {
			handler = conn
		}
	}
	if handler != nil {
		return handler
	}
	if m.notFound == nil {
		return http.HandlerFunc(http.NotFound)
	}
	return m.notFound
}

func (m *Path) handle(prefix string, buf []byte) {
	m.trie.Put([]byte(prefix), buf)
	if m.prefixLength < len(prefix) {
		m.prefixLength = len(prefix)
	}
}

func (m *Path) setHandler(hand http.Handler, reg *regexp.Regexp) ([]byte, error) {
	k := atomic.AddUint32(&m.size, 1)
	buf := make([]byte, handlerMapPrefixSize)
	binary.BigEndian.PutUint32(buf, k)

	_, ok := m.handlers[k]
	if !ok {
		m.handlers[k] = &regexpRoutes{}
	}
	if reg == nil {
		if m.handlers[k].handler != nil {
			return nil, ErrRouteAlreadyExists
		}
		m.handlers[k].handler = hand
	} else {
		m.handlers[k].matchers = append(m.handlers[k].matchers, &matcher{
			match:   reg,
			handler: hand,
		})
	}

	return buf, nil
}

func (m *Path) getHandler(index []byte) (http.Handler, bool) {
	c, ok := m.handlers[binary.BigEndian.Uint32(index)]
	if !ok {
		return nil, false
	}
	if len(c.matchers) != 0 {
		for _, r := range c.matchers {
			if r.match.Match(index[handlerMapPrefixSize:]) {
				return r.handler, true
			}
		}
	}
	if c.handler == nil {
		return nil, false
	}
	return c.handler, true
}

func (m *Path) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	m.Handler(r.URL.Path).ServeHTTP(rw, r)
}
