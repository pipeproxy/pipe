package bzip2

import (
	"github.com/wzshiming/pipe/codec"
	"github.com/wzshiming/pipe/configure"
)

const (
	name = "bzip2"
)

func init() {
	configure.Register(name, NewDecodeWithConfig)
}

func NewDecodeWithConfig() (codec.Decoder, error) {
	return NewCoder(), nil
}
