package config_dump

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/wzshiming/pipe"
	"github.com/wzshiming/pipe/internal/stream"
)

func NewConfigDump(ro bool) *ConfigDump {
	return &ConfigDump{
		readOnly: ro,
	}
}

type ConfigDump struct {
	readOnly bool
}

func (c *ConfigDump) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	default:
		http.NotFound(rw, r)
	case http.MethodGet:
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
	case http.MethodPut:
		if c.readOnly {
			http.NotFound(rw, r)
			return
		}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
		ctx := r.Context()
		pip, ok := pipe.GetPipeWithContext(ctx)
		if !ok {
			http.Error(rw, "bad context", http.StatusBadGateway)
			return
		}

		err = pip.Reload(body)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
		rw.WriteHeader(http.StatusOK)
		stream.CloseExcess()
	}
}
