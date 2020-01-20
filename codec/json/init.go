package json

import (
	"github.com/wzshiming/pipe/codec"
	"github.com/wzshiming/pipe/configure"
)

const name = "json"

func init() {
	configure.Register(name, NewDecoder)
	configure.Register(name, NewEncoder)
}

func NewEncoder() codec.Encoder {
	return NewCoder(nil)
}

func NewDecoder() codec.Decoder {
	return NewCoder(nil)
}
