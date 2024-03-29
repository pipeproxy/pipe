package quic

import (
	"fmt"
	"net/http"

	"github.com/pipeproxy/pipe/components/common/register"
	"github.com/pipeproxy/pipe/components/packet"
	"github.com/pipeproxy/pipe/components/tls"
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
	return NewServer(conf.Handler, conf.TLS), nil
}
