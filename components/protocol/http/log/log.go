package log

import (
	"fmt"
	"net/http"
	"sync/atomic"

	"github.com/felixge/httpsnoop"
	"github.com/pipeproxy/pipe/components/stdio/output"
	"github.com/pipeproxy/pipe/internal/log"
	"github.com/wzshiming/logger"
)

type Log struct {
	handler http.Handler
	output  output.Output
	counter uint64
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

	counter := atomic.AddUint64(&l.counter, 1)
	ll.WithName(fmt.Sprintf("request-%d", counter)).
		Info(u,
			"host", r.Host,
			"method", r.Method,
			"contentLength", r.ContentLength,
			"header", r.Header,
		)
	metric := httpsnoop.CaptureMetrics(l.handler, rw, r)
	ll.WithName(fmt.Sprintf("response-%d", counter)).
		Info(u,
			"code", metric.Code,
			"duration", metric.Duration,
			"contentLength", metric.Written,
			"header", rw.Header(),
		)
}
