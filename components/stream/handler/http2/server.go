package http

import (
	"context"
	"net/http"

	"github.com/wzshiming/pipe/components/stream"
	"github.com/wzshiming/pipe/components/tls"
	"github.com/wzshiming/pipe/internal/listener"
	"github.com/wzshiming/pipe/internal/logger"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type server struct {
	handler   http.Handler
	tlsConfig tls.TLS
}

func NewServer(handler http.Handler, tlsConfig tls.TLS) *server {
	s := &server{
		handler:   handler,
		tlsConfig: tlsConfig,
	}
	return s
}

var h2 = http2.Server{}

func (s *server) serve(ctx context.Context, listen stream.StreamListener, handler http.Handler) error {
	baseContext := func(stream.StreamListener) context.Context {
		return ctx
	}

	svc := http.Server{
		Handler:     handler,
		BaseContext: baseContext,
		TLSConfig:   s.tlsConfig.TLS(),
	}

	svc.Handler = h2c.NewHandler(svc.Handler, &h2)

	err := http2.ConfigureServer(&svc, &h2)
	if err != nil {
		return err
	}

	if svc.TLSConfig != nil {
		err = svc.ServeTLS(listen, "", "")
	} else {
		err = svc.Serve(listen)
	}
	if err != nil && !listener.IsClosedConnError(err) {
		return err
	}
	return nil
}

func (s *server) ServeStream(ctx context.Context, stm stream.Stream) {
	err := s.serve(ctx, listener.NewSingleConnListener(stm), s.handler)
	if err != nil {
		logger.Errorln("[http2]", err)
		return
	}
}
