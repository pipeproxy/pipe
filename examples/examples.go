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
	debug = bind.DefNetHTTPHandlerConfig{
		Name: "debug",
		Def: bind.MuxNetHTTPHandlerConfig{
			Routes: []bind.MuxNetHTTPHandlerRoute{
				{
					Path: "/",
					Handler: bind.MultiNetHTTPHandlerConfig{
						Multi: []bind.HTTPHandler{
							bind.AddResponseHeaderNetHTTPHandlerConfig{
								Key:   "Content-Type",
								Value: "text/html; charset=utf-8",
							},
							bind.DirectNetHTTPHandlerConfig{
								Code: http.StatusOK,
								Body: bind.InlineIoReaderConfig{
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
					Handler: bind.ExpvarNetHTTPHandler{},
				},
				{
					Prefix:  "/pprof/",
					Handler: bind.PprofNetHTTPHandler{},
				},
				{
					Path:    "/config_dump/",
					Handler: bind.ConfigDumpNetHTTPHandler{},
				},
			},
		},
	}

	ExampleDebug = bind.MultiOnceConfig{
		Multi: []bind.Once{
			bind.ServiceOnceConfig{
				Service: bind.MultiServiceConfig{
					Multi: []bind.Service{
						addrToHTTP(":80", debug, nil),
					},
				},
			},
		},
	}

	ExampleFileServer = bind.MultiOnceConfig{
		Multi: []bind.Once{
			bind.ServiceOnceConfig{
				Service: bind.MultiServiceConfig{
					Multi: []bind.Service{
						bind.DefServiceConfig{
							Name: "server",
							Def: addrToHTTP(":80", bind.FileNetHTTPHandlerConfig{
								Root: "",
							}, nil),
						},
					},
				},
			},
		},
	}

	ExampleForward = bind.SampleOnceConfig{
		Pipe: bind.MultiServiceConfig{
			Multi: []bind.Service{
				bind.RefServiceConfig{Name: "host1"},
				bind.RefServiceConfig{Name: "host2"},
				bind.RefServiceConfig{Name: "server"},
			},
		},
		Components: []bind.Component{
			bind.DefNetHTTPHandlerConfig{
				Name: "balance",
				Def: bind.PollerNetHTTPHandlerConfig{
					Poller: bind.PollerNetHTTPHandlerPollerEnumEnumRoundRobin,
					Handlers: []bind.HTTPHandler{
						bind.ForwardNetHTTPHandlerConfig{
							URL: "http://127.0.0.1:8001",
						},
						bind.ForwardNetHTTPHandlerConfig{
							URL: "http://127.0.0.1:8002",
						},
					},
				},
			},
			bind.DefServiceConfig{
				Name: "server",
				Def: addrToHTTP(":80", bind.RefNetHTTPHandlerConfig{
					Name: "balance",
				}, nil),
			},
			bind.DefNetHTTPHandlerConfig{
				Name: "page1",
				Def: bind.DirectNetHTTPHandlerConfig{
					Code: http.StatusOK,
					Body: bind.InlineIoReaderConfig{
						Data: `<html><body>This is Pipe page1 {{.Scheme}}://{{.Host}}{{.RequestURI}}</body></html>`,
					},
				},
			},
			bind.DefServiceConfig{
				Name: "host1",
				Def: addrToHTTP(":8001", bind.RefNetHTTPHandlerConfig{
					Name: "page1",
				}, nil),
			},
			bind.DefNetHTTPHandlerConfig{
				Name: "page2",
				Def: bind.DirectNetHTTPHandlerConfig{
					Code: http.StatusOK,
					Body: bind.InlineIoReaderConfig{
						Data: `<html><body>This is Pipe page2 {{.Scheme}}://{{.Host}}{{.RequestURI}}</body></html>`,
					},
				},
			},
			bind.DefServiceConfig{
				Name: "host2",
				Def: addrToHTTP(":8002", bind.RefNetHTTPHandlerConfig{
					Name: "page2",
				}, nil),
			},
		},
	}

	ExampleHTTPS = bind.SampleOnceConfig{
		Pipe: bind.MultiServiceConfig{
			Multi: []bind.Service{
				bind.RefServiceConfig{
					Name: "server",
				},
			},
		},
		Components: []bind.Component{
			bind.DefNetHTTPHandlerConfig{
				Name: "redirect",
				Def: bind.RedirectNetHTTPHandlerConfig{
					Code:     http.StatusFound,
					Location: "{{.Scheme}}s://{{.Host}}{{.RequestURI}}",
				},
			},

			bind.DefNetHTTPHandlerConfig{
				Name: "page",
				Def: bind.DirectNetHTTPHandlerConfig{
					Code: http.StatusOK,
					Body: bind.InlineIoReaderConfig{
						Data: `<html><body>This is Pipe page {{.Scheme}}://{{.Host}}{{.RequestURI}}</body></html>`,
					},
				},
			},
			bind.DefServiceConfig{
				Name: "server",
				Def: bind.MultiServiceConfig{
					Multi: []bind.Service{
						addrToHTTP(":80", bind.RefNetHTTPHandlerConfig{
							Name: "redirect",
						}, nil),
						addrToHTTP(":443", bind.RefNetHTTPHandlerConfig{
							Name: "page",
						}, bind.SelfSignedTLS{}),
					},
				},
			},
		},
	}

	ExampleWeighted = bind.SampleOnceConfig{
		Pipe: bind.MultiServiceConfig{
			Multi: []bind.Service{
				bind.RefServiceConfig{
					Name: "gateway",
				},
			},
		},
		Components: []bind.Component{
			bind.DefNetHTTPHandlerConfig{
				Name: "weighted",
				Def: bind.WeightedNetHTTPHandlerConfig{
					Weighted: []bind.WeightedNetHTTPHandlerWeighted{
						{
							Weight: 2,
							Handler: bind.DirectNetHTTPHandlerConfig{
								Code: http.StatusOK,
								Body: bind.InlineIoReaderConfig{
									Data: `<html><body>This is Pipe page1 {{.Scheme}}://{{.Host}}{{.RequestURI}}</body></html>`,
								},
							},
						},
						{
							Weight: 8,
							Handler: bind.DirectNetHTTPHandlerConfig{
								Code: http.StatusOK,
								Body: bind.InlineIoReaderConfig{
									Data: `<html><body>This is Pipe page2 {{.Scheme}}://{{.Host}}{{.RequestURI}}</body></html>`,
								},
							},
						},
					},
				},
			},
			bind.DefServiceConfig{
				Name: "gateway",
				Def: addrToHTTP(":80", bind.RefNetHTTPHandlerConfig{
					Name: "weighted",
				}, nil),
			},
		},
	}

	ExampleBasic = bind.SampleOnceConfig{
		Pipe: bind.MultiServiceConfig{
			Multi: []bind.Service{
				bind.RefServiceConfig{
					Name: "server",
				},
			},
		},
		Components: []bind.Component{
			debug,
			bind.DefServiceConfig{
				Name: "server",
				Def: bind.MultiServiceConfig{
					Multi: []bind.Service{
						addrToHTTP(":80", bind.RefNetHTTPHandlerConfig{
							Name: "debug",
						}, nil),
						addrToHTTP(":443", bind.RefNetHTTPHandlerConfig{
							Name: "debug",
						}, bind.SelfSignedTLS{}),
					},
				},
			},
		},
	}
)

func addrToHTTP(address string, handler bind.HTTPHandler, tls bind.TLS) bind.Service {
	return bind.StreamServiceConfig{
		Listener: bind.NetworkStreamListenerListenConfigConfig{
			Network: bind.NetworkStreamListenerListenConfigNetworkEnumEnumTCP,
			Address: address,
		},
		Handler: bind.HTTPStreamHandlerConfig{
			TLS: tls,
			Handler: bind.LogNetHTTPHandlerConfig{
				Output: bind.FileIoWriterConfig{
					Path: "/dev/stderr",
				},
				Handler: handler,
			},
		},
	}
}
