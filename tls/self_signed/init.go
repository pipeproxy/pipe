package self_signed

import (
	"github.com/wzshiming/pipe/configure"
	"github.com/wzshiming/pipe/tls"
)

const name = "self_signed"

func init() {
	configure.Register(name, NewSelfSignedWithConfig)
}

func NewSelfSignedWithConfig() (tls.TLS, error) {
	tlsConfig, err := NewSelfSigned()
	if err != nil {
		return nil, err
	}
	return tls.NewConfig(tlsConfig), nil
}
