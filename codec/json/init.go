package json

import (
	"context"

	"github.com/wzshiming/pipe/codec"
	"github.com/wzshiming/pipe/configure"
)

const name = "json"

func init() {
	configure.Register(name, NewDecoder)
	configure.Register(name, NewEncoder)
}

func NewEncoder(ctx context.Context) (codec.Encoder, error) {
	return NewCoder(nil), nil
}

func NewDecoder(ctx context.Context) (codec.Decoder, error) {
	return NewCoder(nil), nil
}
