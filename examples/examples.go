package examples

import (
	"net/http"

	"github.com/wzshiming/pipe/bind"
)

var Examples = map[string]interface{}{
	"../pipe":        ExampleBasic,
	"web_debug":      ExampleDebug,
	"web_fileserver": ExampleFileServer,
	"web_forward":    ExampleForward,
	"web_https":      ExampleHTTPS,
	"web_weighted":   ExampleWeighted,
}

var (
	debug = bind.NameHTTPHandler{
		Name: "debug",
		HTTPHandler: bind.HTTPHandlerMuxConfig{
			Routes: []bind.HTTPHandlerMuxRoute{
				{
					Path: "/",
					Handler: bind.HTTPHandlerMultiConfig{
						Multi: []bind.HTTPHandler{
							bind.HTTPHandlerAddResponseHeaderConfig{
								Key:   "Content-Type",
								Value: "text/html; charset=utf-8",
							},
							bind.HTTPHandlerDirectConfig{
								Code: http.StatusOK,
								Body: bind.InputInlineConfig{
									Data: `<pre>
<a href="./expvar/">./expvar/</a>
<a href="./pprof/">./pprof/</a>
<a href="./config_dump/">./config_dump/</a>
</pre>`,
								},
							},
						},
					},
				},
				{
					Path:    "/expvar/",
					Handler: bind.HTTPHandlerExpvar{},
				},
				{
					Prefix:  "/pprof/",
					Handler: bind.HTTPHandlerPprof{},
				},
				{
					Path:    "/config_dump/",
					Handler: bind.HTTPHandlerConfigDump{},
				},
			},
		},
	}

	ExampleDebug = bind.OnceConfigConfig{
		Pipe: bind.RefService("server"),
		Components: []bind.PipeComponent{
			bind.NameService{
				Name:    "server",
				Service: addrToHTTP(":80", bind.RefHTTPHandler("debug"), nil),
			},
			debug,
		},
	}

	ExampleFileServer = bind.OnceConfigConfig{
		Pipe: bind.RefService("server"),
		Components: []bind.PipeComponent{
			bind.NameService{
				Name:    "server",
				Service: addrToHTTP(":80", bind.RefHTTPHandler("file"), nil),
			},
			bind.NameHTTPHandler{
				Name: "file",
				HTTPHandler: bind.HTTPHandlerFileConfig{
					Root: "",
				},
			},
		},
	}

	ExampleForward = bind.OnceConfigConfig{
		Pipe: bind.ServiceMultiConfig{
			Multi: []bind.Service{
				bind.RefService("host1"),
				bind.RefService("host2"),
				bind.RefService("gateway"),
			},
		},
		Components: []bind.PipeComponent{
			bind.NameService{
				Name:    "server",
				Service: addrToHTTP(":80", bind.RefHTTPHandler("balance"), nil),
			},
			bind.NameHTTPHandler{
				Name: "balance",
				HTTPHandler: bind.HTTPHandlerPollerConfig{
					Poller: bind.HTTPHandlerPollerPollerEnumEnumRoundRobin,
					Handlers: []bind.HTTPHandler{
						bind.HTTPHandlerForwardConfig{
							URL: "http://127.0.0.1:8001",
						},
						bind.HTTPHandlerForwardConfig{
							URL: "http://127.0.0.1:8002",
						},
					},
				},
			},
			bind.NameService{
				Name:    "host1",
				Service: addrToHTTP(":8001", bind.RefHTTPHandler("page1"), nil),
			},
			bind.NameHTTPHandler{
				Name: "page1",
				HTTPHandler: bind.HTTPHandlerDirectConfig{
					Code: http.StatusOK,
					Body: bind.InputInlineConfig{
						Data: `<html><body>This is Pipe page1 {{.Scheme}}://{{.Host}}{{.RequestURI}}</body></html>`,
					},
				},
			},
			bind.NameService{
				Name:    "host2",
				Service: addrToHTTP(":8002", bind.RefHTTPHandler("page2"), nil),
			},
			bind.NameHTTPHandler{
				Name: "page2",
				HTTPHandler: bind.HTTPHandlerDirectConfig{
					Code: http.StatusOK,
					Body: bind.InputInlineConfig{
						Data: `<html><body>This is Pipe page2 {{.Scheme}}://{{.Host}}{{.RequestURI}}</body></html>`,
					},
				},
			},
		},
	}

	ExampleHTTPS = bind.OnceConfigConfig{
		Pipe: bind.ServiceMultiConfig{
			Multi: []bind.Service{
				bind.RefService("server"),
			},
		},
		Components: []bind.PipeComponent{
			bind.NameService{
				Name: "server",
				Service: bind.ServiceMultiConfig{
					Multi: []bind.Service{
						addrToHTTP(":80", bind.RefHTTPHandler("redirect"), nil),
						addrToHTTP(":443", bind.RefHTTPHandler("page"), bind.TLSSelfSigned{}),
					},
				},
			},

			bind.NameHTTPHandler{
				Name: "redirect",
				HTTPHandler: bind.HTTPHandlerRedirectConfig{
					Code:     http.StatusFound,
					Location: "{{.Scheme}}s://{{.Host}}{{.RequestURI}}",
				},
			},

			bind.NameHTTPHandler{
				Name: "page",
				HTTPHandler: bind.HTTPHandlerDirectConfig{
					Code: http.StatusOK,
					Body: bind.InputInlineConfig{
						Data: `<html><body>This is Pipe page {{.Scheme}}://{{.Host}}{{.RequestURI}}</body></html>`,
					},
				},
			},
		},
	}

	ExampleWeighted = bind.OnceConfigConfig{
		Pipe: bind.ServiceMultiConfig{
			Multi: []bind.Service{
				bind.RefService("gateway"),
			},
		},
		Components: []bind.PipeComponent{
			bind.NameService{
				Name:    "gateway",
				Service: addrToHTTP(":80", bind.RefHTTPHandler("weighted"), nil),
			},
			bind.NameHTTPHandler{
				Name: "weighted",
				HTTPHandler: bind.HTTPHandlerWeightedConfig{
					Weighted: []bind.HTTPHandlerWeightedWeighted{
						{
							Weight: 2,
							Handler: bind.HTTPHandlerDirectConfig{
								Code: http.StatusOK,
								Body: bind.InputInlineConfig{
									Data: `<html><body>This is Pipe page1 {{.Scheme}}://{{.Host}}{{.RequestURI}}</body></html>`,
								},
							},
						},
						{
							Weight: 8,
							Handler: bind.HTTPHandlerDirectConfig{
								Code: http.StatusOK,
								Body: bind.InputInlineConfig{
									Data: `<html><body>This is Pipe page2 {{.Scheme}}://{{.Host}}{{.RequestURI}}</body></html>`,
								},
							},
						},
					},
				},
			},
		},
	}

	ExampleBasic = bind.OnceConfigConfig{
		Pipe: bind.ServiceMultiConfig{
			Multi: []bind.Service{
				bind.RefService("server"),
			},
		},
		Components: []bind.PipeComponent{
			bind.NameService{
				Name: "server",
				Service: bind.ServiceMultiConfig{
					Multi: []bind.Service{
						addrToHTTP(":80", bind.RefHTTPHandler("debug"), nil),
						addrToHTTP(":443", bind.RefHTTPHandler("debug"), bind.TLSSelfSigned{}),
					},
				},
			},
			debug,
		},
	}
)

func addrToHTTP(address string, handler bind.HTTPHandler, tls bind.TLS) bind.Service {
	return bind.ServiceStreamConfig{
		Listener: bind.StreamListenConfigNetworkConfig{
			Network: bind.StreamListenConfigNetworkNetworkEnumEnumTCP,
			Address: address,
		},
		Handler: bind.StreamHandlerHTTPConfig{
			TLS: tls,
			Handler: bind.HTTPHandlerLogConfig{
				Output: bind.OutputFileConfig{
					Path: "/dev/stderr",
				},
				Handler: handler,
			},
		},
	}
}
