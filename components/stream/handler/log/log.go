package log

import (
	"context"
	"fmt"
	"sync/atomic"

	"github.com/pipeproxy/pipe/components/stdio/output"
	"github.com/pipeproxy/pipe/components/stream"
	"github.com/pipeproxy/pipe/internal/log"
	"github.com/wzshiming/logger"
)

type Log struct {
	handler stream.Handler
	output  output.Output
	size    uint64
}

func NewLog(h stream.Handler, o output.Output) *Log {
	return &Log{handler: h, output: o, size: 0}
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

	ll = ll.WithName(fmt.Sprintf("stream-%d", atomic.AddUint64(&l.size, 1)))
	ll = ll.WithValues(
		"localAddress", stm.LocalAddr(),
		"remoteAddress", stm.RemoteAddr(),
	)
	ll.Info("Connect")
	l.handler.ServeStream(logger.WithContext(ctx, ll), stm)
	ll.Info("Disconnect")
}
