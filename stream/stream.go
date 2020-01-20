package stream

import (
	"context"
	"net"
)

type Stream = net.Conn

type Handler interface {
	ServeStream(ctx context.Context, stm Stream)
}
