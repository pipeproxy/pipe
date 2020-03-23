package bzip2

import (
	"github.com/wzshiming/pipe/components/codec"
	"github.com/wzshiming/pipe/components/common/register"
)

const (
	name = "bzip2"
)

func init() {
	register.Register(name, NewDecodeWithConfig)
}

func NewDecodeWithConfig() (codec.Decoder, error) {
	return NewCoder(), nil
}
