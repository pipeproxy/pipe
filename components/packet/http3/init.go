package quic

import (
	"fmt"
	"net/http"

	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/components/packet"
	"github.com/wzshiming/pipe/components/tls"
)

var (
	ErrNotHandler = fmt.Errorf("not handler")
	ErrNotTLS     = fmt.Errorf("not tls")
)

func init() {
	register.Register("http3", NewServerWithConfig)
}

type Config struct {
	Handler http.Handler
	TLS     tls.TLS
}

func NewServerWithConfig(conf *Config) (packet.Handler, error) {
	if conf.Handler == nil {
		return nil, ErrNotHandler
	}
	if conf.TLS == nil {
		return nil, ErrNotTLS
	}
	return NewServer(conf.Handler, conf.TLS.TLS()), nil
}
