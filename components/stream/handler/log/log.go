package log

import (
	"context"
	"fmt"
	"sync/atomic"

	"github.com/pipeproxy/pipe/components/stream"
	"github.com/wzshiming/logger"
)

type Log struct {
	handler stream.Handler
	size    uint64
}

func NewLog(h stream.Handler) *Log {
	return &Log{handler: h, size: 0}
}

func (l *Log) ServeStream(ctx context.Context, stm stream.Stream) {
	log := logger.FromContext(ctx)
	log = log.WithName(fmt.Sprintf("stream-%d", atomic.AddUint64(&l.size, 1)))
	ctx = logger.WithContext(ctx, log)
	log = log.WithValues(
		"localAddress", stm.LocalAddr(),
		"remoteAddress", stm.RemoteAddr(),
	)
	log.Info("Connect")
	l.handler.ServeStream(ctx, stm)
	log.Info("Disconnect")
}
