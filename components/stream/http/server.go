package http

import (
	"context"
	"crypto/tls"
	"net/http"
	"sync"

	"github.com/wzshiming/pipe/components/stream"
	"github.com/wzshiming/pipe/components/stream/listener"
	"github.com/wzshiming/pipe/internal/logger"
)

type server struct {
	handler http.Handler
	tls     *tls.Config
	pool    sync.Pool
}

func NewServer(handler http.Handler, tls *tls.Config) *server {
	s := &server{
		handler: handler,
		tls:     tls,
	}
	return s
}

func (s *server) serve(ctx context.Context, listen listener.StreamListener, handler http.Handler) error {
	baseContext := func(listener.StreamListener) context.Context {
		return ctx
	}
	if s.tls == nil {
		var svc = http.Server{
			Handler:     handler,
			BaseContext: baseContext,
		}

		err := svc.Serve(listen)
		if err != nil && err != http.ErrServerClosed {
			return err
		}
	} else {
		tls, ok := s.pool.Get().(*tls.Config)
		if !ok {
			tls = s.tls.Clone()
		}
		defer s.pool.Put(tls)
		var svc = http.Server{
			Handler:     handler,
			TLSConfig:   tls,
			BaseContext: baseContext,
		}

		err := svc.ServeTLS(listen, "", "")
		if err != nil && err != http.ErrServerClosed {
			return err
		}
	}
	return nil
}

func (s *server) ServeStream(ctx context.Context, stm stream.Stream) {
	err := s.serve(ctx, newSingleConnListener(stm), s.handler)
	if err != nil {
		logger.Error("[http]", err)
		return
	}
}
