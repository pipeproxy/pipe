package log

import (
	"net/http"

	"github.com/felixge/httpsnoop"
	"github.com/pipeproxy/pipe/components/stdio/output"
	"github.com/pipeproxy/pipe/internal/log"
	"github.com/wzshiming/logger"
)

type Log struct {
	handler http.Handler
	output  output.Output
}

func NewLog(h http.Handler, o output.Output) *Log {
	return &Log{handler: h, output: o}
}

func (l *Log) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ll := logger.FromContext(ctx)
	if !ll.Enabled() {
		l.handler.ServeHTTP(rw, r)
		return
	}

	if l.output != nil {
		ll = log.WithOut(ll, l.output)
	}

	u := r.RequestURI

	ll.WithName("Request").
		Info(u,
			"host", r.Host,
			"method", r.Method,
			"contentLength", r.ContentLength,
			"header", r.Header,
		)
	metric := httpsnoop.CaptureMetrics(l.handler, rw, r)
	ll.WithName("Response").
		Info(u,
			"code", metric.Code,
			"duration", metric.Duration,
			"contentLength", metric.Written,
			"header", rw.Header(),
		)
}
