package quit

import (
	"net/http"

	"github.com/pipeproxy/pipe"
	"github.com/pipeproxy/pipe/internal/logger"
)

type Quit struct {
}

func NewQuit() http.Handler {
	return &Quit{}
}

func (c *Quit) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	svc, ok := pipe.GetPipeWithContext(r.Context())
	if !ok {
		http.Error(rw, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
		return
	}
	rw.WriteHeader(http.StatusOK)

	err := svc.Close()
	if err != nil {
		logger.Errorf("service close error: %s", err)
		return
	}
}
