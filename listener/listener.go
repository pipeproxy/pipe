package listener

import (
	"context"
	"net"
)

type ListenConfig interface {
	Listen(ctx context.Context) (net.Listener, error)
}
