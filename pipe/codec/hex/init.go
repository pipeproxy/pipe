package hex

import (
	"github.com/wzshiming/pipe/configure/manager"
	"github.com/wzshiming/pipe/pipe/codec"
)

const (
	name = "hex"
)

func init() {
	manager.Register(name, NewEncodeWithConfig)
	manager.Register(name, NewDecodeWithConfig)
}

func NewEncodeWithConfig() (codec.Encoder, error) {
	return NewCoder(), nil
}

func NewDecodeWithConfig() (codec.Decoder, error) {
	return NewCoder(), nil
}
