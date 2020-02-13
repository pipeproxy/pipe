package acme

import (
	"github.com/wzshiming/pipe/configure/manager"
	"github.com/wzshiming/pipe/pipe/tls"
)

const name = "acme"

func init() {
	manager.Register(name, NewAcmeWithConfig)
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
