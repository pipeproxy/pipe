package codec

import (
	"io"

	"github.com/wzshiming/pipe/configure/alias"
	"github.com/wzshiming/pipe/pipe/common/load"
)

func init() {
	var encoder Encoder
	alias.Register("codec.Encoder", &encoder)
	load.Register(&encoder)
	var decoder Decoder
	alias.Register("codec.Decoder", &decoder)
	load.Register(&decoder)
	var unmarshaler Unmarshaler
	alias.Register("codec.Unmarshaler", &unmarshaler)
	load.Register(&unmarshaler)
	var marshaler Marshaler
	alias.Register("codec.Marshaler", &marshaler)
	load.Register(&marshaler)
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
