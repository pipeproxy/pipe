package network

import (
	"context"
	"log"
	"net"

	"github.com/wzshiming/pipe/decode"
	"github.com/wzshiming/pipe/listener"
)

var ListenConfig net.ListenConfig

func NewNetwork(ctx context.Context, name string, config []byte) (listener.Listener, error) {
	var conf Config
	err := decode.Decode(ctx, config, &conf)
	if err != nil {
		return nil, err
	}
	log.Printf("[INFO] Listen to %s://%s", conf.Network, conf.Address)
	return ListenConfig.Listen(ctx, conf.Network, conf.Address)
}
