package log

import (
	"io"

	"github.com/go-logr/logr"
	"github.com/wzshiming/logger"
	"github.com/wzshiming/logger/zap"
)

var loggerFiles = map[string]logr.Logger{}

func init() {
	l, _ := NewLogFromFile("stderr")
	logger.SetLogger(l)
}

func NewLogFromFile(file string) (logr.Logger, error) {
	l, ok := loggerFiles[file]
	if ok {
		return l, nil
	}

	f, err := NewFile(file)
	if err != nil {
		return nil, err
	}

	l = NewLog(f)
	loggerFiles[file] = l
	return l, nil
}

func NewLog(w io.Writer) logr.Logger {
	return zap.New(zap.WriteTo(w), zap.UseDevMode(true))
}

func WithOut(log logr.Logger, w io.Writer) logr.Logger {
	return zap.WithOut(log, w)
}
