package stream

import (
	"context"
	"sync"
	"time"

	"github.com/wzshiming/pipe/components/stream"
	"github.com/wzshiming/pipe/internal/listener"
	"github.com/wzshiming/pipe/internal/logger"
)

type Server struct {
	listenConfig      stream.ListenConfig
	listener          stream.StreamListener
	handler           stream.Handler
	pool              sync.Map
	disconnectOnClose bool
}

func NewServer(listenConfig stream.ListenConfig, handler stream.Handler, disconnectOnClose bool) (*Server, error) {
	s := &Server{
		listenConfig:      listenConfig,
		handler:           handler,
		disconnectOnClose: disconnectOnClose,
	}

	return s, nil
}

func (s *Server) Run(ctx context.Context) error {
	listen, err := s.listenConfig.ListenStream(ctx)
	if err != nil {
		return err
	}
	s.listener = listen
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			if listener.IsClosedConnError(err) || err == context.Canceled {
				return nil
			}
			logger.Error(err)
			continue
		}
		go s.ServeStream(ctx, conn)
	}
}

func (s *Server) Close() error {
	if s.listener == nil {
		return nil
	}
	err := s.listener.Close()

	if s.disconnectOnClose {
		deadline := time.Now().Add(-1 * time.Minute)
		s.pool.Range(func(key, value interface{}) bool {
			stm := key.(stream.Stream)
			err := stm.SetDeadline(deadline)
			if err != nil {
				addr := stm.LocalAddr()
				logger.Errorf("SetDeadline %s://%s error: %s", addr.Network(), addr.String(), err)
			}
			return true
		})
	}
	return err
}

func (s *Server) ServeStream(ctx context.Context, stm stream.Stream) {
	if s.disconnectOnClose {
		s.pool.Store(stm, ctx)
		defer s.pool.Delete(stm)
	}
	s.handler.ServeStream(ctx, nopCloser{stm})
	err := stm.Close()
	if err != nil {
		addr := stm.LocalAddr()
		logger.Errorf("Close %s://%s error: %s", addr.Network(), addr.String(), err)
		return
	}
}

type nopCloser struct {
	stream.Stream
}

func (nopCloser) Close() error {
	return nil
}
