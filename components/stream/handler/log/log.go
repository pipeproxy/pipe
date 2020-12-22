package log

import (
	"context"
	"fmt"
	"sync/atomic"

	svc_stream "github.com/pipeproxy/pipe/components/service/stream"
	"github.com/pipeproxy/pipe/components/stdio/output"
	"github.com/pipeproxy/pipe/components/stream"
	"github.com/pipeproxy/pipe/internal/log"
	"github.com/wzshiming/logger"
)

type Log struct {
	handler             stream.Handler
	output              output.Output
	originalDestination bool
	counter             uint64
}

func NewLog(h stream.Handler, o output.Output, originalDestination bool) *Log {
	return &Log{handler: h, output: o, originalDestination: originalDestination, counter: 0}
}

func (l *Log) ServeStream(ctx context.Context, stm stream.Stream) {
	ll := logger.FromContext(ctx)
	if !ll.Enabled() {
		l.handler.ServeStream(ctx, stm)
		return
	}

	if l.output != nil {
		ll = log.WithOut(ll, l.output)
	}

	ll = ll.WithName(fmt.Sprintf("stream-%d", atomic.AddUint64(&l.counter, 1)))
	ll = ll.WithValues(
		"localAddress", stm.LocalAddr(),
		"remoteAddress", stm.RemoteAddr(),
	)
	if l.originalDestination {
		if _, d, ok := svc_stream.GetRawStreamAndOriginalDestinationAddrWithContext(ctx); ok {
			ll = ll.WithValues(
				"originalDestinationAddress", d,
			)
		}
	}
	ll.Info("connect")
	l.handler.ServeStream(logger.WithContext(ctx, ll), stm)
	ll.Info("disconnect")
}
