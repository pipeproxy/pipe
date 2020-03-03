package yaml

import (
	"github.com/wzshiming/pipe/configure/decode"
	"github.com/wzshiming/pipe/pipe/codec"
)

const name = "yaml"

func init() {
	decode.Register(name, NewUnmarshaler)
	decode.Register(name, NewMarshaler)
}

func NewMarshaler() codec.Marshaler {
	return NewCoder()
}

func NewUnmarshaler() codec.Unmarshaler {
	return NewCoder()
}
