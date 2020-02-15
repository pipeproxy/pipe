package self_signed

import (
	"github.com/wzshiming/pipe/configure/decode"
	"github.com/wzshiming/pipe/pipe/tls"
)

const name = "self_signed"

func init() {
	decode.Register(name, NewSelfSignedWithConfig)
}

func NewSelfSignedWithConfig() (tls.TLS, error) {
	tlsConfig, err := NewSelfSigned()
	if err != nil {
		return nil, err
	}
	return tls.WrapTLS(tlsConfig), nil
}
