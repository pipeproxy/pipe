package codec

import (
	"io"

	"github.com/wzshiming/pipe/components/common/types"
)

func init() {
	var encoder Encoder
	types.Register(&encoder)
	var decoder Decoder
	types.Register(&decoder)
	var unmarshaler Unmarshaler
	types.Register(&unmarshaler)
	var marshaler Marshaler
	types.Register(&marshaler)
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
