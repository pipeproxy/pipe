package quic

import (
	"context"

	quic "github.com/lucas-clemente/quic-go"
	"github.com/pipeproxy/pipe/components/packet"
	"github.com/pipeproxy/pipe/components/stream"
	"github.com/pipeproxy/pipe/components/tls"
	"github.com/wzshiming/logger"
)

type server struct {
	pkt       packet.Packet
	tlsConfig tls.TLS
}

func NewServer(pkt packet.Packet, tlsConfig tls.TLS) stream.ListenConfig {
	s := &server{
		pkt:       pkt,
		tlsConfig: tlsConfig,
	}
	return s
}

func (s *server) ListenStream(ctx context.Context) (stream.StreamListener, error) {
	log := logger.FromContext(ctx)
	log = log.WithName("quic")
	ctx = logger.WithContext(ctx, log)
	listen, err := quic.Listen(s.pkt, s.tlsConfig.TLS(), &quic.Config{})
	if err != nil {
		return nil, err
	}
	return NewListener(ctx, listen), nil
}
