package http

import (
	"context"
	"crypto/tls"
	"io"
	"log"
	"net"
	"net/http"

	"github.com/wzshiming/pipe"
	"github.com/wzshiming/pipe/stream"
)

type server struct {
	handler http.Handler
	tls     *tls.Config
}

func NewServer(handler http.Handler, tls *tls.Config) *server {
	s := &server{
		handler: handler,
		tls:     tls,
	}
	return s
}

func (s *server) serve(ctx context.Context, listener net.Listener) {
	var svc = http.Server{
		Handler:   s,
		TLSConfig: s.tls,
		BaseContext: func(net.Listener) context.Context {
			return ctx
		},
	}
	if s.tls == nil {
		err := svc.Serve(listener)
		if err != nil && err != io.ErrClosedPipe {
			log.Println("[ERROR] [http]", err)
		}
	} else {
		err := svc.ServeTLS(listener, "", "")
		if err != nil && err != io.ErrClosedPipe {
			log.Println("[ERROR] [http]", err)
		}
	}
}

func (s *server) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Server", pipe.Name)
	s.handler.ServeHTTP(rw, r)
}

func (s *server) ServeStream(ctx context.Context, stm stream.Stream) {
	s.serve(ctx, &singleConnListener{stm})
}
