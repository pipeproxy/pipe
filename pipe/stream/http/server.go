package http

import (
	"context"
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/wzshiming/pipe/pipe/stream"
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

func (s *server) serve(ctx context.Context, listener net.Listener, handler http.Handler) error {
	if s.tls == nil {
		var svc = http.Server{
			Handler: handler,
			BaseContext: func(net.Listener) context.Context {
				return ctx
			},
		}
		err := svc.Serve(listener)
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
			Handler:   handler,
			TLSConfig: tls,
			BaseContext: func(net.Listener) context.Context {
				return ctx
			},
		}
		err := svc.ServeTLS(listener, "", "")
		if err != nil && err != http.ErrServerClosed {
			return err
		}
	}
	return nil
}

func (s *server) ServeStream(ctx context.Context, stm stream.Stream) {
	err := s.serve(ctx, newSingleConnListener(stm), s.handler)
	if err != nil {
		log.Println("[ERROR] [http]", err)
		return
	}
}
