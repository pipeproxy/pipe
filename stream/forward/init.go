package forward

import (
	"github.com/wzshiming/pipe/stream"
)

func init() {
	stream.Register(name, NewForwardWithConfig)
}
