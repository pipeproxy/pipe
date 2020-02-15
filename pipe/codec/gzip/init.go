package gzip

import (
	"github.com/wzshiming/pipe/configure/decode"
	"github.com/wzshiming/pipe/pipe/codec"
)

const (
	name = "gzip"
)

func init() {
	decode.Register(name, NewEncodeWithConfig)
	decode.Register(name, NewDecodeWithConfig)
}

func NewEncodeWithConfig() (codec.Encoder, error) {
	return NewCoder(), nil
}

func NewDecodeWithConfig() (codec.Decoder, error) {
	return NewCoder(), nil
}
