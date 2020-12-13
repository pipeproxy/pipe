package stream

import (
	"context"
	"net"
	"sync"
	"time"

	"github.com/pipeproxy/pipe/components/stream"
	"github.com/pipeproxy/pipe/internal/listener"
	"github.com/wzshiming/logger"
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
	log := logger.FromContext(ctx)
	log = log.WithName("stream")
	ctx = logger.WithContext(ctx, log)
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
			log.Error(err, "listener accept")
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
				logger.Log.Error(err, "SetDeadline",
					"address", addr.String(),
				)
			}
			return true
		})
	}
	return err
}

func (s *Server) ServeStream(ctx context.Context, stm stream.Stream) {
	ctx = withContext(ctx, stm, !s.listenConfig.IsVirtual())
	if s.disconnectOnClose {
		s.pool.Store(stm, ctx)
		defer s.pool.Delete(stm)
	}
	s.handler.ServeStream(ctx, nopCloser{stm})
	err := stm.Close()
	if listener.IsClosedConnError(err) {
		err = nil
	}
	if err != nil {
		logger.FromContext(ctx).Error(err, "close listen")
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

type contextData struct {
	raw         stream.Stream
	originDst   net.Addr
	isOriginDst bool
}

func GetRawStreamWithContext(ctx context.Context) (stream.Stream, bool) {
	i := ctx.Value(streamCtxKeyType(0))
	if i == nil {
		return nil, false
	}
	p, ok := i.(*contextData)
	if !ok || p.raw == nil {
		return nil, false
	}
	return p.raw, true
}

func GetRawStreamAndOriginalDestinationAddrWithContext(ctx context.Context) (stream.Stream, net.Addr, bool) {
	i := ctx.Value(streamCtxKeyType(0))
	if i == nil {
		return nil, nil, false
	}
	p, ok := i.(*contextData)
	if !ok || p.raw == nil || !p.isOriginDst {
		return nil, nil, false
	}
	if p.originDst != nil {
		return p.raw, p.originDst, true
	}
	addr, err := listener.GetOriginalDestinationAddr(p.raw)
	if err != nil {
		logger.FromContext(ctx).Error(err, "GetOriginalDestinationAddr")
		p.isOriginDst = false
		return nil, nil, false
	}
	p.originDst = addr
	return p.raw, addr, true
}

func withContext(ctx context.Context, s stream.Stream, d bool) context.Context {
	return context.WithValue(ctx, streamCtxKeyType(0), &contextData{
		raw:         s,
		isOriginDst: d,
	})
}
