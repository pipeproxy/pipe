package examples

import (
	"net/http"

	"github.com/pipeproxy/pipe/bind"
	"github.com/pipeproxy/pipe/config"
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
	ExampleDebug = config.BuildSampleWithOnce(
		bind.RefServiceConfig{
			Name: "server",
		}, bind.DefServiceConfig{
			Name: "server",
			Def: bind.MultiServiceConfig{
				Multi: []bind.Service{
					config.BuildH1WithService(":8088",
						config.BuildHTTPLogStderr(config.BuildAdminWithHTTPHandler())),
				},
			},
		})

	ExampleFileServer = bind.MultiOnceConfig{
		Multi: []bind.Once{
			bind.ServiceOnceConfig{
				Service: bind.MultiServiceConfig{
					Multi: []bind.Service{
						bind.DefServiceConfig{
							Name: "server",
							Def: config.BuildH1WithService(":80",
								config.BuildHTTPLogStderr(bind.FileNetHTTPHandlerConfig{
									Root: "",
								})),
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
		bind.DefServiceConfig{
			Name: "server",
			Def: config.BuildH1WithService(":80",
				config.BuildHTTPLogStderr(bind.RefNetHTTPHandlerConfig{
					Name: "balance",
				})),
		},
		bind.DefNetHTTPHandlerConfig{
			Name: "balance",
			Def: bind.LbNetHTTPHandlerConfig{
				Policy: bind.LbNetHTTPHandlerLoadBalancePolicyEnumEnumRoundRobin,
				Handlers: []bind.LbNetHTTPHandlerWeight{
					{
						Handler: bind.ForwardNetHTTPHandlerConfig{
							URL: "http://127.0.0.1:8001",
						},
					},
					{
						Handler: bind.ForwardNetHTTPHandlerConfig{
							URL: "http://127.0.0.1:8002",
						},
					},
				},
			},
		},
		bind.DefServiceConfig{
			Name: "host1",
			Def: config.BuildH1WithService(":8001",
				bind.DirectNetHTTPHandlerConfig{
					Code: http.StatusOK,
					Body: bind.InlineIoReaderConfig{
						Data: `<html><body>This is Pipe page1 {{.Scheme}}://{{.Host}}{{.RequestURI}}</body></html>`,
					},
				}),
		},
		bind.DefServiceConfig{
			Name: "host2",
			Def: config.BuildH1WithService(":8002",
				bind.DirectNetHTTPHandlerConfig{
					Code: http.StatusOK,
					Body: bind.InlineIoReaderConfig{
						Data: `<html><body>This is Pipe page2 {{.Scheme}}://{{.Host}}{{.RequestURI}}</body></html>`,
					},
				}),
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
					config.BuildHTTPRedirectToHTTPSWithService(":80"),
					config.BuildH2WithService(":443",
						config.BuildHTTPLogStderr(bind.RefNetHTTPHandlerConfig{
							Name: "page",
						}), bind.SelfSignedTLS{}),
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
			Def: bind.LbNetHTTPHandlerConfig{
				Handlers: []bind.LbNetHTTPHandlerWeight{
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
			Def: config.BuildH1WithService(":80",
				config.BuildHTTPLogStderr(
					bind.RefNetHTTPHandlerConfig{
						Name: "weighted",
					}),
			),
		},
	)

	ExampleBasic = config.BuildSampleWithOnce(
		bind.RefServiceConfig{
			Name: "server",
		},
		bind.DefNetHTTPHandlerConfig{
			Name: "page",
			Def:  config.BuildHomeWithHTTPHandler(),
		},

		bind.DefTLSConfig{
			Name: "tls",
			Def:  bind.SelfSignedTLS{},
		},
		bind.DefServiceConfig{
			Name: "server",
			Def: bind.MultiServiceConfig{
				Multi: []bind.Service{
					config.BuildH1WithService(":80",
						config.BuildHTTPLogStderr(bind.RefNetHTTPHandlerConfig{
							Name: "page",
						}),
					),
					config.BuildH2WithService(":443",
						config.BuildHTTPLogStderr(bind.RefNetHTTPHandlerConfig{
							Name: "page",
						}), bind.RefTLSConfig{
							Name: "tls",
						}),
					config.BuildH3WithService(":443",
						config.BuildHTTPLogStderr(bind.RefNetHTTPHandlerConfig{
							Name: "page",
						}), bind.RefTLSConfig{
							Name: "tls",
						}),
				},
			},
		},
	)
)
