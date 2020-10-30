package validation

import (
	"io/ioutil"

	"github.com/pipeproxy/pipe/components/common/register"
	"github.com/pipeproxy/pipe/components/stdio/input"
	"github.com/pipeproxy/pipe/components/tls"
)

const (
	name = "validation"
)

func init() {
	register.Register(name, NewValidationWithConfig)
}

type Config struct {
	Ca input.Input
}

func NewValidationWithConfig(conf *Config) (tls.TLS, error) {
	ca, err := ioutil.ReadAll(conf.Ca)
	if err != nil {
		return nil, err
	}

	tlsConfig, err := NewValidation(ca)
	if err != nil {
		return nil, err
	}
	return tls.WrapTLS(tlsConfig), nil
}
