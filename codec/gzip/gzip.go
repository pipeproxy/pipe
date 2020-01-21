package gzip

import (
	"compress/gzip"
	"io"
)

type Coder struct {
}

func NewCoder() Coder {
	return Coder{}
}

func (d Coder) Decode(r io.Reader) (io.Reader, error) {
	return gzip.NewReader(r)
}

func (d Coder) Encode(w io.Writer) (io.Writer, error) {
	return gzip.NewWriter(w), nil
}
