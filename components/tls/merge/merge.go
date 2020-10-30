package merge

import (
	"fmt"

	"github.com/pipeproxy/pipe/components/tls"
)

var (
	ErrNotTls = fmt.Errorf("not tls")
)

type Merge struct {
	config []tls.TLS
	res    *tls.Config
}

func NewMerge(config []tls.TLS) *Merge {
	return &Merge{
		config: config,
	}
}

func (m *Merge) TLS() *tls.Config {
	if m.res != nil {
		return m.res
	}
	switch len(m.config) {
	case 0:
		return nil
	case 1:
		return m.config[0].TLS()
	}

	n := &tls.Config{}
	for _, t := range m.config {
		v := t.TLS()
		if v == nil {
			continue
		}
		if v.RootCAs != nil && n.RootCAs == nil {
			n.RootCAs = v.RootCAs
		}
		if v.ClientCAs != nil && n.ClientCAs == nil {
			n.ClientCAs = v.ClientCAs
		}
		if v.ServerName != "" && n.ServerName == "" {
			n.ServerName = v.ServerName
		}
		if v.InsecureSkipVerify && !n.InsecureSkipVerify {
			n.InsecureSkipVerify = v.InsecureSkipVerify
		}
		if len(v.Certificates) != 0 {
			n.Certificates = append(n.Certificates, v.Certificates...)
		}
	}
	m.res = n
	return n
}
