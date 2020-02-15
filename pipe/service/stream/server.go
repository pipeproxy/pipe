package stream

import (
	"context"
	"io"
	"log"
	"net"

	"github.com/wzshiming/pipe/pipe/stream"
	"github.com/wzshiming/pipe/pipe/stream/listener"
)

type Server struct {
	listenConfig listener.ListenConfig
	listener     net.Listener
	handler      stream.Handler
}

func NewServer(listenConfig listener.ListenConfig, handler stream.Handler) (*Server, error) {
	s := &Server{
		listenConfig: listenConfig,
		handler:      handler,
	}
	listener, err := s.listenConfig.ListenStream(context.Background())
	if err != nil {
		return nil, err
	}

	s.listener = listener
	return s, nil
}

func (s *Server) Run(ctx context.Context) error {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			if err == io.ErrClosedPipe {
				return nil
			}
			return err
		}
		go s.ServeStream(ctx, conn)
	}
}

func (s *Server) Close() error {
	return s.listener.Close()
}

func (s *Server) ServeStream(ctx context.Context, stm stream.Stream) {
	s.handler.ServeStream(ctx, nopCloser{stm})
	err := stm.Close()
	if err != nil {
		addr := stm.LocalAddr()
		log.Printf("[ERROR] Close %s://%s error: %s", addr.Network(), addr.String(), err.Error())
		return
	}
}

type nopCloser struct {
	stream.Stream
}

func (nopCloser) Close() error {
	return nil
}
