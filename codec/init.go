package codec

import (
	"github.com/wzshiming/pipe/decode"
)

func init() {
	decode.Register(NewDecoder)
	decode.Register(NewEncoder)
}
