package json

import (
	"github.com/wzshiming/pipe/codec"
)

func init() {
	codec.RegisterDecoder("json", NewDecoder)
	codec.RegisterEncoder("json", NewEncoder)
}
