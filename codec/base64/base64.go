package base64

import (
	"encoding/base64"
	"fmt"
	"io"
)

var (
	ErrNotEncoding = fmt.Errorf("not encoding")
)

type Coder struct {
	enc *base64.Encoding
}

func NewCoder(enc *base64.Encoding) Coder {
	return Coder{enc: enc}
}

func (d Coder) Decode(r io.Reader) (io.Reader, error) {
	return base64.NewDecoder(d.enc, r), nil
}

func (d Coder) Encode(w io.Writer) (io.Writer, error) {
	return base64.NewEncoder(d.enc, w), nil
}
