package bzip2

import (
	"github.com/wzshiming/pipe/codec"
	"github.com/wzshiming/pipe/configure/manager"
)

const (
	name = "bzip2"
)

func init() {
	manager.Register(name, NewDecodeWithConfig)
}

func NewDecodeWithConfig() (codec.Decoder, error) {
	return NewCoder(), nil
}
