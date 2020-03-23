package stream

import (
	"context"
	"net"

	"github.com/wzshiming/pipe/components/common/types"
)

func init() {
	var handler Handler
	types.Register(&handler)
}

type Stream = net.Conn

type Handler interface {
	ServeStream(ctx context.Context, stm Stream)
}
