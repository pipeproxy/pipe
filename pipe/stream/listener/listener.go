package listener

import (
	"context"
	"net"

	"github.com/wzshiming/pipe/configure/alias"
)

func init() {
	var listenConfig ListenConfig
	alias.Register("stream.ListenConfig", &listenConfig)
}

type StreamListener = net.Listener

type ListenConfig interface {
	ListenStream(ctx context.Context) (StreamListener, error)
}
