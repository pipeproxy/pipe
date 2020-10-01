package from

import (
	"crypto/tls"
)

func NewFrom(domain string, cert, key []byte) (*tls.Config, error) {
	pair, err := tls.X509KeyPair(cert, key)
	if err != nil {
		return nil, err
	}
	conf := &tls.Config{}
	conf.ServerName = domain
	conf.InsecureSkipVerify = domain == ""
	conf.Certificates = append(conf.Certificates, pair)
	return conf, nil
}
