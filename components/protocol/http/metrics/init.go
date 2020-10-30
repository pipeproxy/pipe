package metrics

import (
	"github.com/pipeproxy/pipe/components/common/register"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	name = "metrics"
)

func init() {
	register.Register(name, promhttp.Handler)
}
