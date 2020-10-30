package stream

import (
	"context"
	"net"

	"github.com/pipeproxy/pipe/components/common/types"
)

func init() {
	var handler Handler
	types.Register(&handler)
	var listenConfig ListenConfig
	types.Register(&listenConfig)
	var dialer Dialer
	types.Register(&dialer)
	var stream Stream
	types.Register(&stream)
}

type Stream = net.Conn

type StreamListener = net.Listener

type Handler interface {
	ServeStream(ctx context.Context, stm Stream)
}

type ListenConfig interface {
	ListenStream(ctx context.Context) (StreamListener, error)
}

type Dialer interface {
	DialStream(ctx context.Context) (Stream, error)
}
