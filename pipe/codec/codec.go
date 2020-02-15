package codec

import (
	"io"

	"github.com/wzshiming/pipe/configure/alias"
)

func init() {
	var encoder Encoder
	alias.Register("codec.Encoder", &encoder)
	var decoder Decoder
	alias.Register("codec.Decoder", &decoder)
	var unmarshaler Unmarshaler
	alias.Register("codec.Unmarshaler", &unmarshaler)
	var marshaler Marshaler
	alias.Register("codec.Marshaler", &marshaler)
}

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
