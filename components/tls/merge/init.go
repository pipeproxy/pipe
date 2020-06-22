package merge

import (
	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/components/tls"
)

const name = "merge"

func init() {
	register.Register(name, NewMergeWithConfig)
}

type Config struct {
	Merge []tls.TLS
}

func NewMergeWithConfig(conf *Config) (tls.TLS, error) {
	ts := make([]*tls.Config, 0, len(conf.Merge))
	for _, v := range conf.Merge {
		ts = append(ts, v.TLS())
	}
	return tls.WrapTLS(NewMerge(ts)), nil
}
