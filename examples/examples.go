package examples

import (
	"net/http"
	"time"

	"github.com/wzshiming/pipe/bind"
	"github.com/wzshiming/pipe/config"
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
	ExampleDebug = bind.ServiceOnceConfig{
		Service: bind.MultiServiceConfig{
			Multi: []bind.Service{
				addrToHTTP(":8088", config.BuildAdminWithHTTPHandler(), nil),
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

	ExampleForward = config.BuildSampleWithOnce(
		bind.MultiServiceConfig{
			Multi: []bind.Service{
				bind.RefServiceConfig{Name: "host1"},
				bind.RefServiceConfig{Name: "host2"},
				bind.RefServiceConfig{Name: "server"},
			},
		},
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
	)

	ExampleHTTPS = config.BuildSampleWithOnce(
		bind.MultiServiceConfig{
			Multi: []bind.Service{
				bind.RefServiceConfig{
					Name: "server",
				},
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
						Name: "page",
					}, nil),
					addrToHTTP(":443", bind.RefNetHTTPHandlerConfig{
						Name: "page",
					}, bind.SelfSignedTLS{}),
				},
			},
		},
	)

	ExampleWeighted = config.BuildSampleWithOnce(
		bind.MultiServiceConfig{
			Multi: []bind.Service{
				bind.RefServiceConfig{
					Name: "gateway",
				},
			},
		},
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
	)

	ExampleBasic = config.BuildSampleWithOnce(
		bind.MultiServiceConfig{
			Multi: []bind.Service{
				bind.RefServiceConfig{
					Name: "server",
				},
			},
		},
		bind.DefNetHTTPHandlerConfig{
			Name: "page",
			Def:  config.BuildHomeWithHTTPHandler(),
		},
		bind.DefServiceConfig{
			Name: "server",
			Def: bind.MultiServiceConfig{
				Multi: []bind.Service{
					addrToHTTP(":80", bind.RefNetHTTPHandlerConfig{
						Name: "page",
					}, nil),
					addrToHTTP(":443", bind.RefNetHTTPHandlerConfig{
						Name: "page",
					}, bind.SelfSignedTLS{}),
				},
			},
		},
	)
)

func addrToHTTP(address string, handler bind.HTTPHandler, tls bind.TLS) bind.Service {
	output := bind.FileIoWriterConfig{
		Path: "/dev/stderr",
	}
	listen := bind.NetworkStreamListenerListenConfigConfig{
		Network: bind.NetworkStreamListenerListenConfigNetworkEnumEnumTCP,
		Address: address,
	}
	handle := bind.LogNetHTTPHandlerConfig{
		Output:  output,
		Handler: handler,
	}
	return bind.StreamServiceConfig{
		Listener: listen,
		Handler:  config.BuildHTTP443ToHTTPSWithStreamHandler(handle, tls, time.Second),
	}
}
