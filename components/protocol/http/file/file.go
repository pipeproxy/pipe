package file

import (
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	_ "unsafe"
)

type FileServer struct {
	FileSystem http.FileSystem
	AutoIndex  bool
	Indexes    []string
}

func (f *FileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path

	if url != "/" {
		if !strings.HasSuffix(url, "/") {
			if curl := path.Clean(url); curl != url {
				localRedirect(w, r, curl)
				return
			}
			for _, index := range f.Indexes {
				if strings.HasSuffix(url, index) && url[len(url)-len(index)-1] == '/' {
					localRedirect(w, r, "./")
					return
				}
			}
		} else {
			if curl := path.Clean(url); curl != url[:len(url)-1] {
				localRedirect(w, r, curl+"/")
				return
			}
		}
	}

	file, fi, isDir, err := f.open(url)
	if err != nil {
		msg, code := toHTTPError(err)
		http.Error(w, msg, code)
		return
	}
	defer file.Close()

	if isDir {
		if url[len(url)-1] != '/' {
			localRedirect(w, r, path.Base(url)+"/")
			return
		}
		dirList(w, r, file)
		return
	}

	if url[len(url)-1] == '/' {
		localRedirect(w, r, "../"+path.Base(url))
		return
	}
	http.ServeContent(w, r, url, fi.ModTime(), file)
}

func (f *FileServer) open(name string) (retFile http.File, fi os.FileInfo, isDir bool, retErr error) {
	file, err := f.FileSystem.Open(name)
	if err != nil {
		return nil, nil, false, err
	}

	defer func() {
		if retFile != file {
			file.Close()
		}
	}()

	info, err := file.Stat()
	if err != nil {
		return nil, nil, false, err
	}

	isDir = info.IsDir()
	if isDir {
		for _, index := range f.Indexes {
			file, err := f.FileSystem.Open(filepath.Join(name, index))
			if err != nil {
				continue
			}
			stat, err := file.Stat()
			if err != nil || stat.IsDir() {
				file.Close()
				continue
			}
			return file, stat, isDir, nil
		}
		if !f.AutoIndex {
			return nil, nil, false, os.ErrPermission
		}
	}
	return file, info, isDir, nil
}

//go:linkname dirList net/http.dirList
func dirList(w http.ResponseWriter, r *http.Request, f http.File)

//go:linkname localRedirect net/http.localRedirect
func localRedirect(w http.ResponseWriter, r *http.Request, newPath string)

//go:linkname toHTTPError net/http.toHTTPError
func toHTTPError(err error) (string, int)
