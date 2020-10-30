package stream

import (
	"context"
	"sync"
	"time"

	"github.com/pipeproxy/pipe/components/stream"
	"github.com/pipeproxy/pipe/internal/listener"
	"github.com/pipeproxy/pipe/internal/logger"
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
			logger.Errorln(err)
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
	if listener.IsClosedConnError(err) {
		err = nil
	}
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
	ctx = withContext(ctx, stm)
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

type streamCtxKeyType int

func GetRawStreamWithContext(ctx context.Context) (stream.Stream, bool) {
	i := ctx.Value(streamCtxKeyType(0))
	if i == nil {
		return nil, false
	}
	p, ok := i.(stream.Stream)
	return p, ok
}

func withContext(ctx context.Context, s stream.Stream) context.Context {
	return context.WithValue(ctx, streamCtxKeyType(0), s)
}
