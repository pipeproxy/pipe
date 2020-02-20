package stream

import (
	"context"
	"net"

	"github.com/wzshiming/pipe/configure/alias"
	"github.com/wzshiming/pipe/pipe/common/load"
)

func init() {
	var handler Handler
	alias.Register("stream.Handler", &handler)
	load.Register(&handler)
}

type Stream = net.Conn

type Handler interface {
	ServeStream(ctx context.Context, stm Stream)
}
