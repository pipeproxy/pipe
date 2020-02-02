package http

import (
	"context"
	"crypto/tls"
	"io"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/wzshiming/pipe/stream"
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

func (s *server) serve(ctx context.Context, listener net.Listener, handler http.Handler) {
	if s.tls == nil {
		var svc = http.Server{
			Handler: handler,
			BaseContext: func(net.Listener) context.Context {
				return ctx
			},
		}
		err := svc.Serve(listener)
		if err != nil && err != io.ErrClosedPipe {
			log.Println("[ERROR] [http]", err)
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
		if err != nil && err != io.ErrClosedPipe {
			log.Println("[ERROR] [http]", err)
		}
	}
}

func (s *server) ServeStream(ctx context.Context, stm stream.Stream) {
	done := make(chan struct{})
	s.serve(ctx, &singleConnListener{stm},
		http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			s.handler.ServeHTTP(rw, r)
			close(done)
		}))
	<-done
}
