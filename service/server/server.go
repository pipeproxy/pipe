package server

import (
	"context"

	"github.com/wzshiming/pipe/listener"
	"github.com/wzshiming/pipe/stream"
)

type Server struct {
	listener listener.Listener
	handlers []stream.Handler
}

func NewServer(listener listener.Listener, handlers []stream.Handler) (*Server, error) {
	return &Server{
		listener: listener,
		handlers: handlers,
	}, nil
}

func (s *Server) Reload(handlers []stream.Handler) error {
	s.handlers = handlers
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
	handlers := s.handlers
	for _, handler := range handlers {
		handler.ServeStream(ctx, stm)
	}
	stm.Close()
}
