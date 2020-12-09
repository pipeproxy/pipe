package path

import (
	"encoding/binary"
	"fmt"
	"net/http"
	"regexp"
	"sync/atomic"

	"github.com/pipeproxy/pipe/internal/http/template"
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
func (p *Path) NotFound(handler http.Handler) {
	p.notFound = handler
}

func (p *Path) HandlePrefix(prefix string, handler http.Handler) error {
	buf, err := p.setHandler(handler, nil)
	if err != nil {
		return err
	}
	p.handle(prefix, buf)
	return nil
}

func (p *Path) HandlePrefixAndRegexp(prefix, reg string, handler http.Handler) error {
	r, err := regexp.Compile(reg)
	if err != nil {
		return err
	}
	buf, err := p.setHandler(handler, r)
	if err != nil {
		return err
	}
	p.handle(prefix, buf)
	return nil
}

func (p *Path) HandlePath(path string, handler http.Handler) {
	p.paths[path] = handler
	return
}

// Handler returns most matching handler and prefix bytes data to use for the given reader.
func (p *Path) Handler(path string) (handler http.Handler) {
	handler, ok := p.paths[path]
	if ok {
		return handler
	}
	if p.prefixLength != 0 {
		buf := pool.GetBytes()
		defer pool.PutBytes(buf)
		i := copy(buf, path[:])
		data, _, _ := p.trie.Mapping().Get(buf[:i])
		if len(data) != 0 {
			conn, ok := p.getHandler(data)
			if ok {
				handler = conn
			}
		}
		if handler != nil {
			return handler
		}
	}

	if p.notFound == nil {
		return template.NotFoundHandler
	}
	return p.notFound
}

func (p *Path) handle(prefix string, buf []byte) {
	p.trie.Put([]byte(prefix), buf)
	if p.prefixLength < len(prefix) {
		p.prefixLength = len(prefix)
	}
}

func (p *Path) setHandler(hand http.Handler, reg *regexp.Regexp) ([]byte, error) {
	k := atomic.AddUint32(&p.size, 1)
	buf := make([]byte, handlerMapPrefixSize)
	binary.BigEndian.PutUint32(buf, k)

	_, ok := p.handlers[k]
	if !ok {
		p.handlers[k] = &regexpRoutes{}
	}
	if reg == nil {
		if p.handlers[k].handler != nil {
			return nil, ErrRouteAlreadyExists
		}
		p.handlers[k].handler = hand
	} else {
		p.handlers[k].matchers = append(p.handlers[k].matchers, &matcher{
			match:   reg,
			handler: hand,
		})
	}

	return buf, nil
}

func (p *Path) getHandler(index []byte) (http.Handler, bool) {
	c, ok := p.handlers[binary.BigEndian.Uint32(index)]
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

func (p *Path) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	p.Handler(r.URL.Path).ServeHTTP(rw, r)
}
