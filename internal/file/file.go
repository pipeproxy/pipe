package file

import (
	"io"
	"os"
	"path/filepath"

	"github.com/wzshiming/logger"
)

var outfile = map[string]io.WriteCloser{
	"/dev/stdout": nopCloser{os.Stdout},
	"/dev/stderr": nopCloser{os.Stderr},
	"/dev/null":   emptyWriter{},
	"stdout":      nopCloser{os.Stdout},
	"stderr":      nopCloser{os.Stderr},
	"null":        emptyWriter{},
	"":            emptyWriter{},
}

type fileLogger struct {
	file io.WriteCloser
	path string
}

func NewFile(file string) (io.WriteCloser, error) {
	f, ok := outfile[file]
	if ok {
		return f, nil
	}

	if file != "" {
		abs, err := filepath.Abs(file)
		if err != nil {
			return nil, err
		}
		file = abs
	}

	err := os.MkdirAll(filepath.Dir(file), 0755)
	if err != nil {
		return nil, err
	}
	l := &fileLogger{
		path: file,
	}

	f, err = openAppendFile(l.path)
	if err != nil {
		return nil, err
	}
	l.file = f

	return l, nil
}

func (l *fileLogger) Close() error {
	old := l.file
	logger.Log.V(1).Info("close log file", "path", l.path)
	return old.Close()
}

func (l *fileLogger) Write(p []byte) (n int, err error) {
	if l.file != nil {
		return l.file.Write(p)
	}
	return len(p), nil
}

func openAppendFile(filename string) (*os.File, error) {
	return os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
}
