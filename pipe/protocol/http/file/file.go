package file

import (
	"net/http"
)

type File struct {
	fileServer http.Handler
}

func NewFile(dir string) *File {
	return &File{
		fileServer: http.FileServer(http.Dir(dir)),
	}
}

func (f *File) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	f.fileServer.ServeHTTP(rw, r)
}
