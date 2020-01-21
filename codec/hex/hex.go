package hex

import (
	"encoding/hex"
	"io"
)

type Coder struct {
}

func NewCoder() Coder {
	return Coder{}
}

func (d Coder) Decode(r io.Reader) (io.Reader, error) {
	return hex.NewDecoder(r), nil
}

func (d Coder) Encode(w io.Writer) (io.Writer, error) {
	return hex.NewEncoder(w), nil
}
