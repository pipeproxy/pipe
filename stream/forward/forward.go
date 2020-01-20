package forward

import (
	"context"
	"io"
	"log"
	"net"

	"github.com/wzshiming/pipe/stream"
)

type Forward struct {
	network string
	address string
	dialer  net.Dialer
}

func NewForward(network, address string) *Forward {
	return &Forward{
		network: network,
		address: address,
	}
}

func (f *Forward) ServeStream(ctx context.Context, stm stream.Stream) {
	conn, err := f.dialer.DialContext(ctx, f.network, f.address)
	if err != nil {
		log.Printf("[ERROR] Forward to %s://%s error: %s", f.network, f.address, err.Error())
		return
	}
	defer conn.Close()

	log.Printf("[INFO] Forward to %s://%s", f.network, f.address)

	go io.Copy(stm, conn)
	io.Copy(conn, stm)
}
