package http

import (
	"context"
	"net/http"

	"github.com/pipeproxy/pipe/components/stream"
	"github.com/pipeproxy/pipe/components/tls"
	"github.com/pipeproxy/pipe/internal/listener"
	"github.com/wzshiming/logger"
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

func (s *server) serve(ctx context.Context, listen stream.StreamListener, handler http.Handler) error {
	baseContext := func(stream.StreamListener) context.Context {
		return ctx
	}

	var tlsConfig *tls.Config
	if s.tlsConfig != nil {
		tlsConfig = s.tlsConfig.TLS()
		if tlsConfig != nil {
			tlsConfig = tlsConfig.Clone()
		}
	}

	svc := http.Server{
		Handler:     handler,
		BaseContext: baseContext,
		TLSConfig:   tlsConfig,
	}

	go func() {
		<-ctx.Done()
		svc.Shutdown(context.Background())
	}()
	if svc.TLSConfig != nil {
		svc.TLSConfig.NextProtos = strSliceContainsOrSet(svc.TLSConfig.NextProtos, "h2")
		listen = tls.NewListener(listen, svc.TLSConfig)
	} else {
		var h2 http2.Server
		svc.Handler = h2c.NewHandler(svc.Handler, &h2)
		err := http2.ConfigureServer(&svc, &h2)
		if err != nil {
			return err
		}
	}
	return svc.Serve(listen)
}

func (s *server) ServeStream(ctx context.Context, stm stream.Stream) {
	log := logger.FromContext(ctx)
	if log.Enabled() {
		if s.tlsConfig != nil {
			log = log.WithName("http2")
		} else {
			log = log.WithName("h2c")
		}
		ctx = logger.WithContext(ctx, log)
	}
	err := s.serve(ctx, listener.NewSingleConnListener(stm), s.handler)
	if err != nil && !listener.IsClosedConnError(err) && !listener.IsServerClosedError(err) {
		log.Error(err, "http2 server close")
		return
	}
}

func strSliceContainsOrSet(ss []string, s string) []string {
	for _, v := range ss {
		if v == s {
			return ss
		}
	}
	return append(ss, s)
}
