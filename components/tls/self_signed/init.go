package self_signed

import (
	"github.com/pipeproxy/pipe/components/common/register"
	"github.com/pipeproxy/pipe/components/tls"
)

const (
	name = "self_signed"
)

func init() {
	register.Register(name, NewSelfSignedWithConfig)
}

func NewSelfSignedWithConfig() (tls.TLS, error) {
	tlsConfig, err := NewSelfSigned()
	if err != nil {
		return nil, err
	}
	return tls.WrapTLS(tlsConfig), nil
}
