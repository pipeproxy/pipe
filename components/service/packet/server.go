package packet

import (
	"context"
	"sync"

	"github.com/wzshiming/pipe/components/packet"
	"github.com/wzshiming/pipe/internal/logger"
)

type Server struct {
	listenConfig packet.ListenConfig
	packet       packet.Packet
	handler      packet.Handler
	pool         sync.Map
}

func NewServer(listenConfig packet.ListenConfig, handler packet.Handler) (*Server, error) {
	s := &Server{
		listenConfig: listenConfig,
		handler:      handler,
	}

	return s, nil
}

func (s *Server) Run(ctx context.Context) error {
	pkt, err := s.listenConfig.ListenPacket(ctx)
	if err != nil {
		return err
	}
	s.packet = pkt
	s.ServePacket(ctx, pkt)
	return nil
}

func (s *Server) Close() error {
	if s.packet == nil {
		return nil
	}
	return s.packet.Close()
}

func (s *Server) ServePacket(ctx context.Context, pkt packet.Packet) {
	s.handler.ServePacket(ctx, nopCloser{pkt})
	err := pkt.Close()
	if err != nil {
		addr := pkt.LocalAddr()
		logger.Errorf("Close %s://%s error: %s", addr.Network(), addr.String(), err)
		return
	}
}

type nopCloser struct {
	packet.Packet
}

func (nopCloser) Close() error {
	return nil
}
