package template

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestFormat(t *testing.T) {
	var reqHttp *http.Request
	handleHttp := func(rw http.ResponseWriter, r *http.Request) {
		reqHttp = r
	}

	svcHttp := httptest.NewServer(http.HandlerFunc(handleHttp))
	svcHttp.Client().Get(svcHttp.URL + "/path%20/to?q1=v1&q2=v2")
	svcHttp.Close()
	host := strings.TrimPrefix(svcHttp.URL, "http://")
	hostname, port := splitHostPort(host)

	var reqHttps *http.Request
	handleHttps := func(rw http.ResponseWriter, r *http.Request) {
		reqHttps = r
	}

	svcHttps := httptest.NewTLSServer(http.HandlerFunc(handleHttps))
	svcHttps.Client().Get(svcHttps.URL + "/path%20/to?q1=v1&q2=v2")
	svcHttps.Close()

	type args struct {
		format string
		r      *http.Request
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			args: args{`{{.Scheme}}`, reqHttp},
			want: "http",
		},
		{
			args: args{`{{.Scheme}}`, reqHttps},
			want: "https",
		},
		{
			args: args{`{{.Host}}`, reqHttp},
			want: host,
		},
		{
			args: args{`{{.Hostname}}`, reqHttp},
			want: hostname,
		},
		{
			args: args{`{{.Port}}`, reqHttp},
			want: port,
		},
		{
			args: args{`{{.Path}}`, reqHttp},
			want: "/path /to",
		},
		{
			args: args{`{{.RawPath}}`, reqHttp},
			want: "/path%20/to",
		},
		{
			args: args{`{{.RawQuery}}`, reqHttp},
			want: "q1=v1&q2=v2",
		},
		{
			args: args{`{{.IsQuery}}`, reqHttp},
			want: "?",
		},
		{
			args: args{`{{.RequestURI}}`, reqHttp},
			want: "/path%20/to?q1=v1&q2=v2",
		},
		{
			args: args{`{{.Query "q1"}}`, reqHttp},
			want: "v1",
		},
		{
			args: args{`{{.Query "q2"}}`, reqHttp},
			want: "v2",
		},
		{
			args: args{`{{.Header "Accept-Encoding"}}`, reqHttp},
			want: "gzip",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			temp, err := NewFormat(tt.args.format)
			if err != nil {
				t.Errorf("newTemplate() error = %v", err)
				return
			}
			got, err := temp.FormatString(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Format() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Format() got = %v, want %v", got, tt.want)
			}
		})
	}
}
