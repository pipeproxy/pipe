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

func NewEncoder() (codec.Encoder, error) {
	return NewCoder(nil), nil
}

func NewDecoder() (codec.Decoder, error) {
	return NewCoder(nil), nil
}
