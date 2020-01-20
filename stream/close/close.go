package close

import (
	"context"
	"log"

	"github.com/wzshiming/pipe/stream"
)

type Close struct {
}

func NewClose() *Close {
	return &Close{}
}

func (c *Close) ServeStream(ctx context.Context, stm stream.Stream) {
	err := stm.Close()
	if err != nil {
		addr := stm.LocalAddr()
		log.Printf("[ERROR] Close %s://%s error: %s", addr.Network(), addr.String(), err.Error())
		return
	}
}
