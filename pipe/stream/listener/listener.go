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

type ListenConfig interface {
	ListenStream(ctx context.Context) (net.Listener, error)
}
