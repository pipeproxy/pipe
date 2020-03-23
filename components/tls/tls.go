package tls

type wrapTLS struct {
	tlsConfig *Config
}

func WrapTLS(tlsConfig *Config) TLS {
	return &wrapTLS{tlsConfig: tlsConfig}
}

func (t *wrapTLS) TLS() *Config {
	return t.tlsConfig
}
