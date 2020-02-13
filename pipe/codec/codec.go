package codec

import (
	"io"
)

type Encoder interface {
	Encode(w io.Writer) (io.Writer, error)
}

type Decoder interface {
	Decode(r io.Reader) (io.Reader, error)
}

type Unmarshaler interface {
	Unmarshal(buf []byte, v interface{}) error
}

type Marshaler interface {
	Marshal(v interface{}) ([]byte, error)
}
