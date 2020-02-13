package base64

import (
	"encoding/base32"
	"fmt"
	"io"
)

var (
	ErrNotEncoding = fmt.Errorf("not encoding")
)

type Coder struct {
	enc *base32.Encoding
}

func NewCoder(enc *base32.Encoding) Coder {
	return Coder{enc: enc}
}

func (d Coder) Decode(r io.Reader) (io.Reader, error) {
	return base32.NewDecoder(d.enc, r), nil
}

func (d Coder) Encode(w io.Writer) (io.Writer, error) {
	return base32.NewEncoder(d.enc, w), nil
}
