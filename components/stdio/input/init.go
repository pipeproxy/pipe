package input

import (
	"io"

	"github.com/pipeproxy/pipe/components/common/types"
)

func init() {
	var input Input
	types.Register(&input)
}

type Input = io.Reader

type lazyReader struct {
	reader io.Reader
	err    error
	init   func() (io.Reader, error)
}

func NewLazyReader(f func() (io.Reader, error)) io.Reader {
	return &lazyReader{init: f}
}

func (l *lazyReader) Read(p []byte) (int, error) {
	if l.err != nil {
		return 0, l.err
	}

	if l.reader == nil {
		l.reader, l.err = l.init()
		if l.err != nil {
			return 0, l.err
		}
	}
	return l.reader.Read(p)
}

type readerWithAutoClose struct {
	reader io.Reader
	closer io.Closer
	err    error
}

func NewReaderWithAutoClose(reader io.Reader, closer io.Closer) io.Reader {
	if closer == nil {
		return reader
	}
	return &readerWithAutoClose{
		reader: reader,
		closer: closer,
	}
}

func (l *readerWithAutoClose) Read(p []byte) (int, error) {
	if l.err != nil {
		return 0, l.err
	}
	n, err := l.reader.Read(p)
	if err != nil {
		l.closer.Close()
		l.err = err
		return 0, err
	}
	return n, nil
}
