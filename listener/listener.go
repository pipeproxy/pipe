package listener

import (
	"net"
)

type Listener interface {
	net.Listener
}
