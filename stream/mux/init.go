package mux

import (
	"github.com/wzshiming/pipe/stream"
)

func init() {
	stream.Register(name, NewMuxWithConfig)
}
