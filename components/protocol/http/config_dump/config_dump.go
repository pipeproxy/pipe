package config_dump

import (
	"bytes"
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

	config = bytes.TrimSpace(config)

	contentType := "application/json; charset=utf-8"
	if config[0] != '{' {
		contentType = "text/x-yaml; charset=utf-8"
	}

	rw.Header().Set("Content-Type", contentType)
	rw.Write(config)
}
