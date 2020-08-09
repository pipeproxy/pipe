package quic

import (
	"fmt"

	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/components/packet"
	"github.com/wzshiming/pipe/components/stream"
	"github.com/wzshiming/pipe/components/tls"
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
	return NewServer(conf.Packet, conf.TLS.TLS()), nil
}
