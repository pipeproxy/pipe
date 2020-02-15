package bzip2

import (
	"github.com/wzshiming/pipe/configure/decode"
	"github.com/wzshiming/pipe/pipe/codec"
)

const (
	name = "bzip2"
)

func init() {
	decode.Register(name, NewDecodeWithConfig)
}

func NewDecodeWithConfig() (codec.Decoder, error) {
	return NewCoder(), nil
}
