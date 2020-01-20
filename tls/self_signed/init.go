package self_signed

import (
	"crypto/tls"
	"time"

	"github.com/wzshiming/pipe/configure"
)

const name = "self_signed"

func init() {
	configure.Register(name, NewSelfSigned)
}

func NewSelfSigned() (*tls.Config, error) {
	certBytes, keyBytes, err := GenerateSelfSigned("pipe test", []string{"localhost", "127.0.0.1", "::1"}, time.Now(), time.Hour*24)
	if err != nil {
		return nil, err
	}
	cert, err := tls.X509KeyPair(certBytes, keyBytes)
	if err != nil {
		return nil, err
	}
	conf := &tls.Config{}
	conf.ServerName = "localhost"
	conf.Certificates = append(conf.Certificates, cert)
	return conf, nil
}
