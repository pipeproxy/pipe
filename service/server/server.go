package server

import (
	"context"
	"net"

	"github.com/wzshiming/pipe/stream"
)

type Server struct {
	listener net.Listener
	handler  stream.Handler
}

func NewServer(listener net.Listener, handler stream.Handler) *Server {
	return &Server{
		listener: listener,
		handler:  handler,
	}
}

func (s *Server) Reload(handler stream.Handler) error {
	s.handler = handler
	return nil
}

func (s *Server) Run() error {
	for {
		conn, err := s.listener.Accept()
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
