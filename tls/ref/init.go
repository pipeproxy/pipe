package ref

import (
	"context"
	"crypto/tls"
	"fmt"

	"github.com/wzshiming/pipe/components"
	"github.com/wzshiming/pipe/configure"
)

var (
	ErrNotTlsConfig = fmt.Errorf("not tls config")
)

const name = "ref"

func init() {
	configure.Register(name, NewRefWithConfig)
}

type Config struct {
	Ref string
}

func NewRefWithConfig(ctx context.Context, conf *Config) (*tls.Config, error) {
	components, ok := components.GetCtxComponents(ctx)
	if !ok || components == nil || components.TlsConfigs == nil {
		return nil, ErrNotTlsConfig
	}
	tlsConfig, ok := components.TlsConfigs[conf.Ref]
	if !ok {
		return nil, fmt.Errorf("%s: %w", conf.Ref, ErrNotTlsConfig)
	}
	return tlsConfig, nil
}
