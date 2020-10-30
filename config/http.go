package config

import (
	"fmt"
	"net/http"
	"time"

	"github.com/pipeproxy/pipe/bind"
)

func BuildContentTypeHTMLWithHTTPHandler() bind.HTTPHandler {
	return bind.AddResponseHeaderNetHTTPHandlerConfig{
		Key:   "Content-Type",
		Value: "text/html; charset=utf-8",
	}
}

func BuildHTTPRedirectWithStreamHandler(location string, wait time.Duration) bind.StreamHandler {
	if wait == 0 {
		return bind.HTTP1StreamHandlerConfig{
			Handler: bind.RedirectNetHTTPHandlerConfig{
				Code:     http.StatusFound,
				Location: location,
			},
		}
	}

	return bind.HTTP1StreamHandlerConfig{
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
		return bind.HTTP1StreamHandlerConfig{
			Handler: handler,
		}
	}

	redirect := BuildHTTPRedirectWithStreamHandler("{{.Scheme}}s://{{.Host}}{{.RequestURI}}", wait)

	return bind.MuxStreamHandlerConfig{
		Routes: []bind.MuxStreamHandlerRoute{
			{
				Pattern: "http",
				Handler: redirect,
			},
		},
		NotFound: bind.TLSDownStreamHandlerConfig{
			Handler: bind.HTTP2StreamHandlerConfig{
				Handler: handler,
			},
			TLS: tls,
		},
	}
}

func BuildH3WithService(address string, handler bind.HTTPHandler, tls bind.TLS) bind.Service {
	listen := bind.ListenerPacketListenConfigConfig{
		Network: bind.ListenerPacketListenConfigListenerNetworkEnumEnumUDP,
		Address: address,
	}
	return bind.PacketServiceConfig{
		Listener: listen,
		Handler: bind.HTTP3PacketHandlerConfig{
			Handler: handler,
			TLS:     tls,
		},
	}
}

func BuildH2WithService(address string, handler bind.HTTPHandler, tls bind.TLS) bind.Service {
	listen := bind.ListenerStreamListenConfigConfig{
		Network: bind.ListenerStreamListenConfigListenerNetworkEnumEnumTCP,
		Address: address,
	}
	return bind.StreamServiceConfig{
		Listener: listen,
		Handler: bind.MuxStreamHandlerConfig{
			Routes: []bind.MuxStreamHandlerRoute{
				{
					Pattern: "http",
					Handler: BuildHTTPRedirectWithStreamHandler("{{.Scheme}}s://{{.Host}}{{.RequestURI}}", 0),
				},
			},
			NotFound: bind.HTTP2StreamHandlerConfig{
				Handler: handler,
				TLS:     tls,
			},
		},
	}
}

func BuildH2SupportH3WithService(address string, handler bind.HTTPHandler, tls bind.TLS) bind.Service {
	return BuildH2WithService(address, bind.MultiNetHTTPHandlerConfig{
		Multi: []bind.HTTPHandler{
			bind.AddResponseHeaderNetHTTPHandlerConfig{
				Key:   "Alt-Svc",
				Value: `h3-29=":443"; ma=2592000`,
			},
			handler,
		},
	}, tls)
}

func BuildHTTPRedirectToHTTPSWithService(address string) bind.Service {
	listen := bind.ListenerStreamListenConfigConfig{
		Network: bind.ListenerStreamListenConfigListenerNetworkEnumEnumTCP,
		Address: address,
	}

	return bind.StreamServiceConfig{
		Listener: listen,
		Handler:  BuildHTTPRedirectWithStreamHandler("{{.Scheme}}s://{{.Host}}{{.RequestURI}}", 0),
	}
}

func BuildH1WithService(address string, handler bind.HTTPHandler) bind.Service {
	listen := bind.ListenerStreamListenConfigConfig{
		Network: bind.ListenerStreamListenConfigListenerNetworkEnumEnumTCP,
		Address: address,
	}
	return bind.StreamServiceConfig{
		Listener: listen,
		Handler: bind.HTTP1StreamHandlerConfig{
			Handler: handler,
		},
	}
}

func BuildHTTPLog(log string, handler bind.HTTPHandler) bind.HTTPHandler {
	return bind.LogNetHTTPHandlerConfig{
		Output: bind.FileIoWriterConfig{
			Path: log,
		},
		Handler: handler,
	}
}

func BuildHTTPLogStderr(handler bind.HTTPHandler) bind.HTTPHandler {
	return BuildHTTPLog("/dev/stderr", handler)
}
