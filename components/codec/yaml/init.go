package yaml

import (
	"github.com/wzshiming/pipe/components/codec"
	"github.com/wzshiming/pipe/components/common/register"
)

const name = "yaml"

func init() {
	register.Register(name, NewUnmarshaler)
	register.Register(name, NewMarshaler)
}

func NewMarshaler() codec.Marshaler {
	return NewCoder()
}

func NewUnmarshaler() codec.Unmarshaler {
	return NewCoder()
}
