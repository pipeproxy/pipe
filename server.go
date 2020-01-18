package pipe

import (
	"context"
	"log"

	"github.com/wzshiming/pipe/listener"
	"github.com/wzshiming/pipe/stream"
)

type Server struct {
	Listener listener.Listener
	Handlers []stream.Handler
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

func (s *Server) Start() error {
	go func() {
		err := s.Run()
		if err != nil {
			log.Printf("[ERROR] server start error: %s", err.Error())
		}
	}()
	return nil
}

func (s *Server) ServeStream(ctx context.Context, stm stream.Stream) {
	handlers := s.Handlers
	for _, handler := range handlers {
		handler.ServeStream(ctx, stm)
	}
	stm.Close()
}
