package log

import (
	"net/http"

	"github.com/felixge/httpsnoop"
	"github.com/wzshiming/logger"
)

type Log struct {
	handler http.Handler
}

func NewLog(h http.Handler) *Log {
	return &Log{h}
}

func (l *Log) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	if !log.Enabled() {
		l.handler.ServeHTTP(rw, r)
		return
	}
	u := r.RequestURI

	log.WithName("Request").
		Info(u,
			"host", r.Host,
			"method", r.Method,
			"contentLength", r.ContentLength,
			"header", r.Header,
		)
	metric := httpsnoop.CaptureMetrics(l.handler, rw, r)
	log.WithName("Response").
		Info(u,
			"code", metric.Code,
			"duration", metric.Duration,
			"contentLength", metric.Written,
			"header", rw.Header(),
		)
}
