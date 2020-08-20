package template

import (
	"net/http"
)

type RequestVal struct {
	r *http.Request
}

func (t *RequestVal) Host() string {
	return t.r.Host
}

func (t *RequestVal) Proto() string {
	return t.r.Proto
}

func (t *RequestVal) ProtoMajor() int {
	return t.r.ProtoMajor
}

func (t *RequestVal) ProtoMinor() int {
	return t.r.ProtoMinor
}

func (t *RequestVal) Scheme() string {
	if t.r.TLS == nil {
		return "http"
	}
	return "https"
}

func (t *RequestVal) Path() string {
	return t.r.URL.Path
}

func (t *RequestVal) RawPath() string {
	return t.r.URL.EscapedPath()
}

func (t *RequestVal) Hostname() string {
	host, _ := splitHostPort(t.r.Host)
	return host
}

func (t *RequestVal) Port() string {
	_, port := splitHostPort(t.r.Host)
	return port
}

func (t *RequestVal) RequestURI() string {
	return t.r.RequestURI
}

func (t *RequestVal) RawQuery() string {
	return t.r.URL.RawQuery
}

func (t *RequestVal) IsQuery() string {
	if t.r.URL.ForceQuery || t.r.URL.RawQuery != "" {
		return "?"
	}
	return ""
}

func (t *RequestVal) Query(key string) string {
	return t.r.URL.Query().Get(key)
}

func (t *RequestVal) Header(key string) string {
	return t.r.Header.Get(key)
}

func (t *RequestVal) Method() string {
	return t.r.Method
}

func (t *RequestVal) ContentLength() int64 {
	return t.r.ContentLength
}

func (t *RequestVal) RemoteAddr() string {
	return t.r.RemoteAddr
}

func (t *RequestVal) Form(key string) string {
	if t.r.Form == nil {
		err := t.r.ParseForm()
		if err != nil {
			return ""
		}
	}
	return t.r.Form.Get(key)
}

func (t *RequestVal) PostForm(key string) string {
	if t.r.PostForm == nil {
		err := t.r.ParseForm()
		if err != nil {
			return ""
		}
	}
	return t.r.PostForm.Get(key)
}

func (t *RequestVal) Cookie(key string) string {
	c, err := t.r.Cookie(key)
	if err != nil {
		return ""
	}
	return c.Value
}

func (t *RequestVal) Username() string {
	username, _, _ := t.r.BasicAuth()
	return username
}

func (t *RequestVal) Password() string {
	_, password, _ := t.r.BasicAuth()
	return password
}

func (t *RequestVal) UserAgent() string {
	return t.r.UserAgent()
}
