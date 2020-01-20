package network

import (
	"context"
	"log"
	"net"

	"github.com/wzshiming/pipe/configure"
	"github.com/wzshiming/pipe/listener"
)

func init() {
	configure.Register(name, NewNetworkWithConfig)
}

const name = "network"

type Config struct {
	Network string
	Address string
}

var ListenConfig net.ListenConfig

func NewNetworkWithConfig(ctx context.Context, config []byte) (listener.Listener, error) {
	var conf Config
	err := configure.Decode(ctx, config, &conf)
	if err != nil {
		return nil, err
	}
	log.Printf("[INFO] Listen to %s://%s", conf.Network, conf.Address)
	return ListenConfig.Listen(ctx, conf.Network, conf.Address)
}
