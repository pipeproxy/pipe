package protocol

import (
	"github.com/wzshiming/pipe/decode"
)

func init() {
	decode.Register(NewProtocol)
}
