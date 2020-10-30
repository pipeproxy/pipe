package input

import (
	"io"
	"sync"

	"github.com/pipeproxy/pipe/components/common/types"
)

func init() {
	var input Input
	types.Register(&input)
}

type Input = io.Reader

type LazyReader struct {
	reader Input
	err    error
	init   func() (Input, error)
	once   sync.Once
}

func NewLazyReader(f func() (Input, error)) *LazyReader {
	return &LazyReader{init: f}
}

func (l *LazyReader) Read(p []byte) (int, error) {
	l.once.Do(func() {
		l.reader, l.err = l.init()
	})
	if l.err != nil {
		return 0, l.err
	}
	return l.reader.Read(p)
}
