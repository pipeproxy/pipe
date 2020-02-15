package listener

import (
	"context"
	"net"
)

type ListenConfig interface {
	ListenStream(ctx context.Context) (net.Listener, error)
}
