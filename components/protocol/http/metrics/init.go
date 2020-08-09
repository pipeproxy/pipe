package metrics

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/wzshiming/pipe/components/common/register"
)

const (
	name = "metrics"
)

func init() {
	register.Register(name, promhttp.Handler)
}
