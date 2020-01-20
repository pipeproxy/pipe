package from_file

import (
	"crypto/tls"

	"github.com/wzshiming/pipe/configure"
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

func NewFromFileWithConfig(conf *Config) (*tls.Config, error) {
	return NewFromFile(conf.Domain, conf.CertFile, conf.KeyFile)
}
