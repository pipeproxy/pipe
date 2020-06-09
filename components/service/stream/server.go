package stream

import (
	"context"
	"io"

	"github.com/wzshiming/pipe/components/stream"
	"github.com/wzshiming/pipe/components/stream/listener"
	"github.com/wzshiming/pipe/internal/logger"
)

type Server struct {
	listenConfig listener.ListenConfig
	listener     listener.StreamListener
	handler      stream.Handler
}

func NewServer(listenConfig listener.ListenConfig, handler stream.Handler) (*Server, error) {
	s := &Server{
		listenConfig: listenConfig,
		handler:      handler,
	}
	return s, nil
}

func (s *Server) Run(ctx context.Context) error {
	listener, err := s.listenConfig.ListenStream(ctx)
	if err != nil {
		return err
	}
	s.listener = listener
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			if err == io.ErrClosedPipe || err == context.Canceled {
				return nil
			}
			return err
		}
		go s.ServeStream(ctx, conn)
	}
}

func (s *Server) Close() error {
	if s.listener == nil {
		return nil
	}
	return s.listener.Close()
}

func (s *Server) ServeStream(ctx context.Context, stm stream.Stream) {
	s.handler.ServeStream(ctx, nopCloser{stm})
	err := stm.Close()
	if err != nil {
		addr := stm.LocalAddr()
		logger.Errorf("Close %s://%s error: %s", addr.Network(), addr.String(), err.Error())
		return
	}
}

type nopCloser struct {
	stream.Stream
}

func (nopCloser) Close() error {
	return nil
}
