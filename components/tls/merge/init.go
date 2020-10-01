package merge

import (
	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/components/tls"
)

const (
	name = "merge"
)

func init() {
	register.Register(name, NewMergeWithConfig)
}

type Config struct {
	Merge []tls.TLS
}

func NewMergeWithConfig(conf *Config) (tls.TLS, error) {
	if len(conf.Merge) == 0 {
		return nil, ErrNotTls
	}
	return NewMerge(conf.Merge), nil
}
