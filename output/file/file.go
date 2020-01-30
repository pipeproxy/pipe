package file

import (
	"io"
	"os"
	"path/filepath"
	"sync"

	"github.com/wzshiming/pipe/output"
)

type File struct {
	file io.WriteCloser
	path string
}

var outfile = map[string]output.Output{
	"/dev/stdout": nopCloser{os.Stdout},
	"/dev/stderr": nopCloser{os.Stderr},
	"/dev/null":   nopCloser{emptyWriter{}},
	"":            nopCloser{os.Stderr},
}

type nopCloser struct{ io.Writer }

func (nopCloser) Close() error { return nil }

type emptyWriter struct{}

func (emptyWriter) Write(p []byte) (n int, err error) { return len(p), nil }

var (
	openedFiles = map[string]*File{}
	mut         sync.Mutex
)

func NewFile(path string) (output.Output, error) {
	if l, ok := outfile[path]; ok {
		return l, nil
	}

	mut.Lock()
	defer mut.Unlock()
	if l, ok := openedFiles[path]; ok {
		if ok {
			return l, nil
		}
	}

	file, err := newFile(path)
	if err != nil {
		return nil, err
	}
	openedFiles[path] = file
	return file, nil
}

func newFile(path string) (*File, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	path = abs

	err = os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return nil, err
	}

	file, err := openAppendFile(path)
	if err != nil {
		return nil, err
	}

	return &File{
		path: path,
		file: file,
	}, nil
}

func (f *File) Close() error {
	mut.Lock()
	defer mut.Unlock()
	delete(openedFiles, f.path)
	return f.file.Close()
}

func (f *File) Write(p []byte) (n int, err error) {
	return f.file.Write(p)
}

func openAppendFile(filename string) (*os.File, error) {
	return os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
}
