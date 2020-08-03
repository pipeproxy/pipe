package config

import (
	"fmt"
	"net/http"
	"time"

	"github.com/wzshiming/pipe/bind"
)

func BuildContentTypeHTMLWithHTTPHandler() bind.HTTPHandler {
	return bind.AddResponseHeaderNetHTTPHandlerConfig{
		Key:   "Content-Type",
		Value: "text/html; charset=utf-8",
	}
}

func BuildHTTPRedirectWithHTTPHandler(location string, wait time.Duration) bind.StreamHandler {
	if wait == 0 {
		return bind.HTTPStreamHandlerConfig{
			Handler: bind.RedirectNetHTTPHandlerConfig{
				Code:     http.StatusFound,
				Location: location,
			},
		}
	}

	return bind.HTTPStreamHandlerConfig{
		Handler: bind.MultiNetHTTPHandlerConfig{
			Multi: []bind.HTTPHandler{
				BuildContentTypeHTMLWithHTTPHandler(),
				bind.DirectNetHTTPHandlerConfig{
					Code: http.StatusOK,
					Body: bind.InlineIoReaderConfig{
						Data: fmt.Sprintf(`<meta http-equiv="refresh" content="%.f; url={{.Scheme}}s://{{.Host}}{{.RequestURI}}"/>
<p>Redirect to %s in %s</p>
`, float64(wait)/float64(time.Second), location, wait),
					},
				},
			},
		},
	}
}

func BuildHTTP443ToHTTPSWithStreamHandler(handler bind.HTTPHandler, tls bind.TLS, wait time.Duration) bind.StreamHandler {
	if tls == nil {
		return bind.HTTPStreamHandlerConfig{
			Handler: handler,
		}
	}

	redirect := BuildHTTPRedirectWithHTTPHandler("{{.Scheme}}s://{{.Host}}{{.RequestURI}}", wait)

	return bind.MuxStreamHandlerConfig{
		Routes: []bind.MuxStreamHandlerRoute{
			{
				Pattern: "http",
				Handler: redirect,
			},
		},
		NotFound: bind.HTTPStreamHandlerConfig{
			TLS:     tls,
			Handler: handler,
		},
	}
}
