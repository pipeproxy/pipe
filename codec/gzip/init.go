package gzip

import (
	"github.com/wzshiming/pipe/codec"
	"github.com/wzshiming/pipe/configure"
)

const (
	name = "gzip"
)

func init() {
	configure.Register(name, NewEncodeWithConfig)
	configure.Register(name, NewDecodeWithConfig)
}

func NewEncodeWithConfig() (codec.Encoder, error) {
	return NewCoder(), nil
}

func NewDecodeWithConfig() (codec.Decoder, error) {
	return NewCoder(), nil
}
