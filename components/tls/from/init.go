package from

import (
	"io/ioutil"

	"github.com/pipeproxy/pipe/components/common/register"
	"github.com/pipeproxy/pipe/components/stdio/input"
	"github.com/pipeproxy/pipe/components/tls"
)

const (
	name = "from"
)

func init() {
	register.Register(name, NewFromWithConfig)
}

type Config struct {
	Domain string
	Cert   input.Input
	Key    input.Input
}

func NewFromWithConfig(conf *Config) (tls.TLS, error) {
	cert, err := ioutil.ReadAll(conf.Cert)
	if err != nil {
		return nil, err
	}
	key, err := ioutil.ReadAll(conf.Key)
	if err != nil {
		return nil, err
	}
	tlsConfig, err := NewFrom(conf.Domain, cert, key)
	if err != nil {
		return nil, err
	}
	return tls.WrapTLS(tlsConfig), nil
}
