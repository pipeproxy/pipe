package from_file

import (
	"crypto/tls"
)

func NewFromFile(domain string, certFile string, keyFile string) (*tls.Config, error) {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}
	conf := &tls.Config{}
	conf.ServerName = domain
	conf.Certificates = append(conf.Certificates, cert)
	return conf, nil
}
