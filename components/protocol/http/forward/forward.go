package forward

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"sync"

	"github.com/pipeproxy/pipe/components/stream"
	"github.com/pipeproxy/pipe/internal/http/template"
	"github.com/pipeproxy/pipe/internal/pool"
	"github.com/pipeproxy/pipe/internal/round_tripper"
	"github.com/wzshiming/logger"
	"golang.org/x/net/http2"
)

type Forward struct {
	url       template.Format
	transport http.RoundTripper
	dialer    stream.Dialer
	once      sync.Once
	h2c       bool
}

func NewForward(url string, dialer stream.Dialer, h2c bool) (*Forward, error) {
	f := &Forward{
		dialer: dialer,
		h2c:    h2c,
	}
	if url == "" {
		url = "{{.Scheme}}://{{.Host}}"
	}

	u, err := template.NewFormat(url)
	if err != nil {
		return nil, err
	}
	f.url = u
	return f, nil
}

func (h *Forward) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.once.Do(func() {
		transport := round_tripper.RoundTripper(h.dialer)
		if h.h2c {
			if t, ok := transport.(*http.Transport); ok {
				err := http2.ConfigureTransport(t)
				if err != nil {
					logger.FromContext(r.Context()).Error(err, "http2 ConfigureTransport")
				}
			}
		}
		h.transport = transport
	})
	proxy := httputil.ReverseProxy{
		BufferPool:   pool.Bytes,
		Transport:    h.transport,
		ErrorHandler: errorHandler,
	}
	u := h.url.FormatString(r)
	target, err := url.Parse(u)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	proxy.Director = directorFunc(target)
	proxy.ServeHTTP(rw, r)
}

func directorFunc(target *url.URL) func(req *http.Request) {
	return func(req *http.Request) {
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path = singleJoiningSlash(target.Path, req.URL.Path)
		if target.RawQuery == "" || req.URL.RawQuery == "" {
			req.URL.RawQuery = target.RawQuery + req.URL.RawQuery
		} else {
			req.URL.RawQuery = target.RawQuery + "&" + req.URL.RawQuery
		}
		if _, ok := req.Header["User-Agent"]; !ok {
			// explicitly disable User-Agent so it's not set to default value
			req.Header.Set("User-Agent", "")
		}
	}
}

func errorHandler(rw http.ResponseWriter, r *http.Request, err error) {
	http.Error(rw, err.Error(), http.StatusInternalServerError)
}

func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}
