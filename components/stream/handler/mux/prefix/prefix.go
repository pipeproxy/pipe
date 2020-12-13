package prefix

import (
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"strings"
	"sync/atomic"

	"github.com/pipeproxy/pipe/components/stream"
	"github.com/wzshiming/crun"
	"github.com/wzshiming/logger"
	"github.com/wzshiming/trie"
)

var (
	ErrNotFound = fmt.Errorf("error not found")
)

// Prefix is an Applicative protocol multiplexer
// It matches the prefix of each incoming reader against a list of registered patterns
// and calls the handler for the pattern that most closely matches the Handler.
type Prefix struct {
	trie         *trie.Trie
	prefixLength int
	size         uint32
	handlers     map[uint32]stream.Handler
	notFound     stream.Handler
}

// NewPrefix create a new Prefix.
func NewPrefix() *Prefix {
	p := &Prefix{
		trie:     trie.NewTrie(),
		handlers: map[uint32]stream.Handler{},
	}

	return p
}

// NotFound replies to the handler with an Handler not found error.
func (p *Prefix) NotFound(handler stream.Handler) {
	p.notFound = handler
}

func (p *Prefix) HandleRegexp(pattern string, handler stream.Handler) error {
	if !strings.HasPrefix(pattern, "^") {
		return fmt.Errorf("only prefix matching is supported, change to %q", "^"+pattern)
	}
	r, err := crun.Compile(pattern)
	if err != nil {
		return err
	}

	if size := r.Size(); size > 1000 {
		return fmt.Errorf("regular is too large: %d", size)
	}

	buf := p.setHandler(handler)
	r.Range(func(prefix string) bool {
		p.handle(prefix, buf)
		return true
	})
	return nil
}

func (p *Prefix) HandlePrefix(prefix string, handler stream.Handler) {
	buf := p.setHandler(handler)
	p.handle(prefix, buf)
	return
}

// Handler returns most matching handler and prefix bytes data to use for the given reader.
func (p *Prefix) Handler(r io.Reader) (handler stream.Handler, prefix []byte, err error) {
	if p.prefixLength == 0 {
		return nil, nil, ErrNotFound
	}
	parent := p.trie.Mapping()
	off := 0
	prefix = make([]byte, p.prefixLength)
	for {
		i, err := r.Read(prefix[off:])
		if err != nil {
			return nil, nil, err
		}
		if i == 0 {
			break
		}

		data, next, _ := parent.Get(prefix[off : off+i])
		if len(data) != 0 {
			conn, ok := p.getHandler(data)
			if ok {
				handler = conn
			}
		}

		off += i
		if next == nil {
			break
		}
		parent = next
	}

	if handler == nil {
		if p.notFound == nil {
			return nil, prefix[:off], ErrNotFound
		}
		handler = p.notFound
	}
	return handler, prefix[:off], nil
}

func (p *Prefix) handle(prefix string, buf []byte) {
	p.trie.Put([]byte(prefix), buf)
	if p.prefixLength < len(prefix) {
		p.prefixLength = len(prefix)
	}
}

func (p *Prefix) setHandler(hand stream.Handler) []byte {
	k := atomic.AddUint32(&p.size, 1)
	p.handlers[k] = hand
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, k)
	return buf
}

func (p *Prefix) getHandler(index []byte) (stream.Handler, bool) {
	c, ok := p.handlers[binary.BigEndian.Uint32(index)]
	return c, ok
}

// ServeStream dispatches the reader to the handler whose pattern most closely matches the reader.
func (p *Prefix) ServeStream(ctx context.Context, stm stream.Stream) {
	connector, buf, err := p.Handler(stm)
	if err != nil {
		logger.FromContext(ctx).Error(err, "",
			"prefix", buf,
		)
		stm.Close()
		return
	}
	stm = UnreadStream(stm, buf)
	connector.ServeStream(ctx, stm)
}
