package mux

import (
	"encoding/binary"
	"fmt"
	"net/http"
	"sync/atomic"
	"unsafe"

	"github.com/wzshiming/trie"
)

var (
	ErrNotFound = fmt.Errorf("error not found")
)

// Mux is an Applicative protocol multiplexer.
type Mux struct {
	trie         *trie.Trie
	prefixLength int
	size         uint32
	handlers     map[uint32]http.Handler
	paths        map[string]http.Handler
	notFound     http.Handler
}

// NewMux create a new Mux.
func NewMux() *Mux {
	p := &Mux{
		trie:     trie.NewTrie(),
		handlers: map[uint32]http.Handler{},
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
	buf := m.setHandler(handler)
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
	parent := m.trie.Mapping()
	data, _, _ := parent.Get(*(*[]byte)(unsafe.Pointer(&path)))
	conn, ok := m.getHandler(data)
	if ok {
		handler = conn
	}
	if handler == nil {
		if m.notFound == nil {
			return nil, ErrNotFound
		}
		handler = m.notFound
	}
	return handler, nil
}

func (m *Mux) handle(prefix string, buf []byte) {
	m.trie.Put(*(*[]byte)(unsafe.Pointer(&prefix)), buf)
	if m.prefixLength < len(prefix) {
		m.prefixLength = len(prefix)
	}
}

func (m *Mux) setHandler(hand http.Handler) []byte {
	k := atomic.AddUint32(&m.size, 1)
	m.handlers[k] = hand
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, k)
	return buf
}

func (m *Mux) getHandler(index []byte) (http.Handler, bool) {
	c, ok := m.handlers[binary.BigEndian.Uint32(index)]
	return c, ok
}

func (m *Mux) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	connector, err := m.Handler(path)
	if err != nil {
		http.NotFound(rw, r)
		return
	}
	connector.ServeHTTP(rw, r)
}
