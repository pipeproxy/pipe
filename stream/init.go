package stream

import (
	"github.com/wzshiming/pipe/decode"
)

func init() {
	decode.Register(NewHandler)
}
