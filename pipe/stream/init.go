package stream

import (
	"context"
	"net"

	"github.com/wzshiming/pipe/configure/alias"
)

func init() {
	var handler Handler
	alias.Register("stream.Handler", &handler)
}

type Stream = net.Conn

type Handler interface {
	ServeStream(ctx context.Context, stm Stream)
}
