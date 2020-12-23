package log

import (
	"io"

	"github.com/go-logr/logr"
	"github.com/pipeproxy/pipe/internal/file"
	"github.com/wzshiming/logger"
	"github.com/wzshiming/logger/zap"
)

var loggerFiles = map[string]logr.Logger{}

func init() {
	l, _ := NewLogFromFile("stderr")
	logger.SetLogger(l)
}

func NewLogFromFile(f string) (logr.Logger, error) {
	l, ok := loggerFiles[f]
	if ok {
		return l, nil
	}

	out, err := file.NewFile(f)
	if err != nil {
		return nil, err
	}

	l = NewLog(out)
	loggerFiles[f] = l
	return l, nil
}

func NewLog(w io.Writer) logr.Logger {
	return zap.New(zap.WriteTo(w), zap.UseDevMode(true))
}

func WithOut(log logr.Logger, w io.Writer) logr.Logger {
	return zap.WithOut(log, w)
}
