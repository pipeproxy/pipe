package bzip2

import (
	"compress/bzip2"
	"io"
)

type Coder struct {
}

func NewCoder() Coder {
	return Coder{}
}

func (d Coder) Decode(r io.Reader) (io.Reader, error) {
	return bzip2.NewReader(r), nil
}
