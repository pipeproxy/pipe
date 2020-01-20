package network

import (
	"context"
	"log"
	"net"

	"github.com/wzshiming/pipe/decode"
	"github.com/wzshiming/pipe/listener"
)

func init() {
	decode.Register(name, NewNetworkWithConfig)
}

const name = "network"

type Config struct {
	Network string
	Address string
}

var ListenConfig net.ListenConfig

func NewNetworkWithConfig(ctx context.Context, config []byte) (listener.Listener, error) {
	var conf Config
	err := decode.Decode(ctx, config, &conf)
	if err != nil {
		return nil, err
	}
	log.Printf("[INFO] Listen to %s://%s", conf.Network, conf.Address)
	return ListenConfig.Listen(ctx, conf.Network, conf.Address)
}
