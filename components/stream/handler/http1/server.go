package http

import (
	"context"
	"net/http"

	"github.com/pipeproxy/pipe/components/stream"
	"github.com/pipeproxy/pipe/internal/listener"
	"github.com/wzshiming/logger"
)

type server struct {
	handler http.Handler
}

func NewServer(handler http.Handler) *server {
	s := &server{
		handler: handler,
	}
	return s
}

func (s *server) serve(ctx context.Context, listen stream.StreamListener, handler http.Handler) error {
	baseContext := func(stream.StreamListener) context.Context {
		return ctx
	}

	svc := http.Server{
		Handler:     handler,
		BaseContext: baseContext,
	}

	err := svc.Serve(listen)
	if err != nil && !listener.IsClosedConnError(err) {
		return err
	}
	return nil
}

func (s *server) ServeStream(ctx context.Context, stm stream.Stream) {
	log := logger.FromContext(ctx)
	if log.Enabled() {
		log = log.WithName("http1")
		ctx = logger.WithContext(ctx, log)
	}
	err := s.serve(ctx, listener.NewSingleConnListener(stm), s.handler)
	if err != nil {
		log.Error(err, "http1 server close")
		return
	}
}
