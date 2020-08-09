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
						BuildContentTypeHTMLWithHTTPHandler(),
						bind.DirectNetHTTPHandlerConfig{
							Code: http.StatusOK,
							Body: bind.InlineIoReaderConfig{
								Data: `<pre>
<a href="pprof/">{{.Path}}pprof/</a>
<a href="expvar">{{.Path}}expvar</a>
<a href="must_quit">{{.Path}}must_quit</a>
<a href="healthy">{{.Path}}healthy</a>
<a href="metrics">{{.Path}}metrics</a>
<a href="config_dump">{{.Path}}config_dump</a>
<a href="config_dump_edit.sh">{{.Path}}config_dump_edit.sh</a>
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
			{
				Path:    "/config_dump",
				Handler: bind.ConfigDumpNetHTTPHandlerConfig{},
			},
			{
				Path:    "/metrics",
				Handler: bind.MetricsNetHTTPHandler{},
			},
			{
				Path: "/config_dump_edit.sh",
				Handler: bind.MultiNetHTTPHandlerConfig{
					Multi: []bind.HTTPHandler{
						bind.DirectNetHTTPHandlerConfig{
							Code: http.StatusOK,
							Body: bind.InlineIoReaderConfig{
								Data: `#!/bin/sh
URL="{{.Scheme}}://{{.Host}}"
RESOURCE="$URL/config_dump"
TMP=.pipe_edit_tmp_file.yaml

# Check if editing is allowed
curl -sL -v -X OPTIONS "$RESOURCE" 2>&1 | \
grep "< Allow:" | grep "PUT" > /dev/null || \
{ echo "Editing Not Allowed"; exit 1;}

# Editing
curl -sL "$RESOURCE?yaml" > $TMP && \
vi $TMP && \
curl -sL -X PUT "$RESOURCE" -d "$(cat $TMP)" && \
rm $TMP

# sh -c "$(curl -sL {{.Scheme}}://{{.Host}}{{.Path}})"
`,
							},
						},
					},
				},
			},
		},
	}
}
