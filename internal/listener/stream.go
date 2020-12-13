package listener

import (
	"net"

	"github.com/mikioh/tcp"
)

func GetOriginalDestinationAddr(conn net.Conn) (net.Addr, error) {
	c, err := tcp.NewConn(conn)
	if err != nil {
		return nil, err
	}
	return c.OriginalDst()
}
