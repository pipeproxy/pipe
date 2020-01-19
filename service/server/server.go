package server

import (
	"context"

	"github.com/wzshiming/pipe/decode"
	"github.com/wzshiming/pipe/listener"
	"github.com/wzshiming/pipe/service"
	"github.com/wzshiming/pipe/stream"
)

type Server struct {
	Listener listener.Listener
	Handlers []stream.Handler
}

func NewServerWithConfig(ctx context.Context, name string, config []byte) (service.Service, error) {
	var conf Config
	err := decode.Decode(ctx, config, &conf)
	if err != nil {
		return nil, err
	}
	return NewServer(conf.Listener, conf.Handlers)
}

func NewServer(listener listener.Listener, handlers []stream.Handler) (*Server, error) {
	return &Server{
		Listener: listener,
		Handlers: handlers,
	}, nil
}

func (s *Server) Reload(handlers []stream.Handler) error {
	s.Handlers = handlers
	return nil
}

func (s *Server) Run() error {
	for {
		conn, err := s.Listener.Accept()
		if err != nil {
			return err
		}

		go s.ServeStream(context.Background(), conn)
	}
}

func (s *Server) Close() error {
	return s.Listener.Close()
}

func (s *Server) ServeStream(ctx context.Context, stm stream.Stream) {
	handlers := s.Handlers
	for _, handler := range handlers {
		handler.ServeStream(ctx, stm)
	}
	stm.Close()
}
