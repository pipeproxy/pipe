package quic

import (
	"fmt"

	"github.com/pipeproxy/pipe/components/common/register"
	"github.com/pipeproxy/pipe/components/packet"
	"github.com/pipeproxy/pipe/components/stream"
	"github.com/pipeproxy/pipe/components/tls"
)

var (
	ErrNotPacket = fmt.Errorf("not packet")
	ErrNotTLS    = fmt.Errorf("not tls")
)

func init() {
	register.Register("quic", NewServerWithConfig)
}

type Config struct {
	Packet packet.Packet
	TLS    tls.TLS
}

func NewServerWithConfig(conf *Config) (stream.ListenConfig, error) {
	if conf.Packet == nil {
		return nil, ErrNotPacket
	}
	if conf.TLS == nil {
		return nil, ErrNotTLS
	}
	return NewServer(conf.Packet, conf.TLS), nil
}
