package from_file

import (
	"github.com/wzshiming/pipe/configure"
	"github.com/wzshiming/pipe/tls"
)

const name = "from_file"

func init() {
	configure.Register(name, NewFromFileWithConfig)
}

type Config struct {
	Domain   string
	CertFile string
	KeyFile  string
}

func NewFromFileWithConfig(conf *Config) (tls.TLS, error) {
	tlsConfig, err := NewFromFile(conf.Domain, conf.CertFile, conf.KeyFile)
	if err != nil {
		return nil, err
	}
	return tls.NewConfig(tlsConfig), nil
}
