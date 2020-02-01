package from

import (
	"io/ioutil"

	"github.com/wzshiming/pipe/configure"
	"github.com/wzshiming/pipe/input"
	"github.com/wzshiming/pipe/tls"
)

const name = "from"

func init() {
	configure.Register(name, NewFromWithConfig)
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
	err = conf.Cert.Close()
	if err != nil {
		return nil, err
	}

	key, err := ioutil.ReadAll(conf.Key)
	if err != nil {
		return nil, err
	}
	err = conf.Key.Close()
	if err != nil {
		return nil, err
	}

	tlsConfig, err := NewFrom(conf.Domain, cert, key)
	if err != nil {
		return nil, err
	}
	return tls.NewConfig(tlsConfig), nil
}
