package config

import (
	"net/http"

	"github.com/wzshiming/pipe/bind"
)

func BuildAdminWithHTTPHandler() bind.HTTPHandler {
	return bind.MuxNetHTTPHandlerConfig{
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
{{.Scheme}}://{{.Host}}{{.RequestURI}}
<a href="{{.Scheme}}://{{.Host}}{{.Path}}pprof/">{{.Path}}pprof/</a>
<a href="{{.Scheme}}://{{.Host}}{{.Path}}expvar">{{.Path}}expvar</a>
<a href="{{.Scheme}}://{{.Host}}{{.Path}}config_dump">{{.Path}}config_dump</a>
<a href="{{.Scheme}}://{{.Host}}{{.Path}}must_quit">{{.Path}}must_quit</a>
<a href="{{.Scheme}}://{{.Host}}{{.Path}}healthy">{{.Path}}healthy</a>
</pre>`,
							},
						},
					},
				},
			},
			{
				Prefix:  "/pprof/",
				Handler: bind.PprofNetHTTPHandler{},
			},
			{
				Path:    "/expvar",
				Handler: bind.ExpvarNetHTTPHandler{},
			},
			{
				Path:    "/config_dump",
				Handler: bind.ConfigDumpNetHTTPHandlerConfig{},
			},
			{
				Path:    "/must_quit",
				Handler: bind.QuitNetHTTPHandler{},
			},
			{
				Path: "/healthy",
				Handler: bind.DirectNetHTTPHandlerConfig{
					Code: http.StatusOK,
					Body: bind.InlineIoReaderConfig{
						Data: `healthy`,
					},
				},
			},
		},
	}
}
