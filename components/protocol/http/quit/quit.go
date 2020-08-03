package quit

import (
	"net/http"

	"github.com/wzshiming/pipe"
	"github.com/wzshiming/pipe/internal/logger"
	"github.com/wzshiming/pipe/internal/stream"
)

func NewQuit() *ConfigQuit {
	return &ConfigQuit{}
}

type ConfigQuit struct {
}

func (c *ConfigQuit) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	svc, ok := pipe.GetPipeWithContext(r.Context())
	if !ok {
		http.Error(rw, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
		return
	}
	rw.WriteHeader(http.StatusOK)

	err := svc.Close()
	if svc == nil {
		logger.Errorf("service close error: %s", err)
		return
	}
	stream.CloseExcess()
}
