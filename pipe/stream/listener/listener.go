package listener

import (
	"context"
	"net"

	"github.com/wzshiming/pipe/configure/alias"
	"github.com/wzshiming/pipe/pipe/common/load"
)

func init() {
	var listenConfig ListenConfig
	alias.Register("stream.ListenConfig", &listenConfig)
	load.Register(&listenConfig)
}

type StreamListener = net.Listener

type ListenConfig interface {
	ListenStream(ctx context.Context) (StreamListener, error)
}
