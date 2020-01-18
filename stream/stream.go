package stream

import (
	"context"
	"net"
)

type Stream = net.Conn

type Handler interface {
	ServeStream(ctx context.Context, stm Stream)
}

type HandlerFunc func(ctx context.Context, stm Stream)

func (h HandlerFunc) ServeConn(ctx context.Context, stm Stream) {
	h(ctx, stm)
}
