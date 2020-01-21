package json

import (
	"github.com/wzshiming/pipe/codec"
	"github.com/wzshiming/pipe/configure"
)

const name = "json"

func init() {
	configure.Register(name, NewUnmarshaler)
	configure.Register(name, NewMarshaler)
}

func NewMarshaler() codec.Marshaler {
	return NewCoder()
}

func NewUnmarshaler() codec.Unmarshaler {
	return NewCoder()
}
