package json

import (
	"github.com/wzshiming/pipe/codec"
	"github.com/wzshiming/pipe/configure/manager"
)

const name = "json"

func init() {
	manager.Register(name, NewUnmarshaler)
	manager.Register(name, NewMarshaler)
}

func NewMarshaler() codec.Marshaler {
	return NewCoder()
}

func NewUnmarshaler() codec.Unmarshaler {
	return NewCoder()
}
