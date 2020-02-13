package config_dump

import (
	"net/http"

	"github.com/wzshiming/pipe"
)

type ConfigDump struct {
}

func (c *ConfigDump) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	pip, ok := pipe.GetPipeWithContext(ctx)
	config := []byte("{}")
	if ok {
		config = pip.Config()
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Write(config)
}
