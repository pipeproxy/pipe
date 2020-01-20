package template

import (
	"bytes"
	"io"
	"net/http"
	"strings"
	"sync"
	"text/template"
)

type Template struct {
	template *template.Template
}

func (t *Template) init(text string) error {
	temp := template.New("_")
	temp, err := temp.Parse(text)
	if err != nil {
		return err
	}
	t.template = temp
	return nil
}

func (t *Template) Format(w io.Writer, r *http.Request) error {
	err := t.template.Execute(w, &RequestVal{r})
	if err != nil {
		return err
	}
	return nil
}

func (t *Template) FormatString(r *http.Request) string {
	buf := poolBuffer.Get().(*bytes.Buffer)
	buf.Reset()
	defer poolBuffer.Put(buf)
	err := t.Format(buf, r)
	if err != nil {
		return ""
	}
	return buf.String()
}

var poolBuffer = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 1024))
	},
}

// splitHostPort separates host and port. If the port is not valid, it returns
// the entire input as host, and it doesn't check the validity of the host.
// Unlike net.SplitHostPort, but per RFC 3986, it requires ports to be numeric.
func splitHostPort(hostport string) (host, port string) {
	host = hostport

	colon := strings.LastIndexByte(host, ':')
	if colon != -1 && validOptionalPort(host[colon:]) {
		host, port = host[:colon], host[colon+1:]
	}

	if strings.HasPrefix(host, "[") && strings.HasSuffix(host, "]") {
		host = host[1 : len(host)-1]
	}

	return
}

// validOptionalPort reports whether port is either an empty string
// or matches /^:\d*$/
func validOptionalPort(port string) bool {
	if port == "" {
		return true
	}
	if port[0] != ':' {
		return false
	}
	for _, b := range port[1:] {
		if b < '0' || b > '9' {
			return false
		}
	}
	return true
}
