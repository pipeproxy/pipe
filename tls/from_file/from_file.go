package from_file

import (
	"crypto/tls"
)

func NewFromFile(domain string, cert, key []byte) (*tls.Config, error) {
	pair, err := tls.X509KeyPair(cert, key)
	if err != nil {
		return nil, err
	}
	conf := &tls.Config{}
	conf.ServerName = domain
	conf.Certificates = append(conf.Certificates, pair)
	return conf, nil
}
