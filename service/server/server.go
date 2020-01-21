package server

import (
	"context"
	"net"

	"github.com/wzshiming/pipe/listener"
	"github.com/wzshiming/pipe/stream"
)

type Server struct {
	listenConfig listener.ListenConfig
	listener     net.Listener
	handler      stream.Handler
}

func NewServer(listener listener.ListenConfig, handler stream.Handler) *Server {
	return &Server{
		listenConfig: listener,
		handler:      handler,
	}
}

func (s *Server) Reload(handler stream.Handler) error {
	s.handler = handler
	return nil
}

func (s *Server) Run() error {
	listener, err := s.listenConfig.Listen(context.Background())
	if err != nil {
		return err
	}

	s.listener = listener
	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		go s.ServeStream(context.Background(), conn)
	}
}

func (s *Server) Close() error {
	return s.listener.Close()
}

func (s *Server) ServeStream(ctx context.Context, stm stream.Stream) {
	s.handler.ServeStream(ctx, stm)
}
