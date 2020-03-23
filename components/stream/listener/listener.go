package listener

import (
	"context"
	"net"

	"github.com/wzshiming/pipe/components/common/types"
)

func init() {
	var listenConfig ListenConfig
	types.Register(&listenConfig)
}

type StreamListener = net.Listener

type ListenConfig interface {
	ListenStream(ctx context.Context) (StreamListener, error)
}
