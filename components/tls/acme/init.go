package acme

import (
	"github.com/pipeproxy/pipe/components/common/register"
	"github.com/pipeproxy/pipe/components/tls"
)

const (
	name = "acme"
)

func init() {
	register.Register(name, NewAcmeWithConfig)
}

type Config struct {
	Domains  []string
	CacheDir string
}

func NewAcmeWithConfig(conf *Config) (tls.TLS, error) {
	tlsConfig, err := NewAcme(conf.Domains, conf.CacheDir)
	if err != nil {
		return nil, err
	}
	return tls.WrapTLS(tlsConfig), nil
}
