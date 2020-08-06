package config_dump

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/kubernetes-sigs/yaml"
	"github.com/wzshiming/pipe"
)

func NewConfigDump(ro bool) *ConfigDump {
	return &ConfigDump{
		readOnly: ro,
	}
}

type ConfigDump struct {
	readOnly bool
}

var (
	allowRW = strings.Join([]string{http.MethodHead, http.MethodGet, http.MethodPut}, ",")
	allowRO = strings.Join([]string{http.MethodHead, http.MethodGet}, ",")
)

func (c *ConfigDump) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	default:
		http.Error(rw, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	case http.MethodOptions:
		header := rw.Header()
		if c.readOnly {
			header.Set("Allow", allowRO)
		} else {
			header.Set("Allow", allowRW)
		}
	case http.MethodGet, http.MethodHead:
		contentType := "application/json; charset=utf-8"
		ctx := r.Context()
		pip, ok := pipe.GetPipeWithContext(ctx)
		if !ok {
			http.Error(rw, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
			return
		}

		config := pip.Config()
		if len(r.URL.RawQuery) >= 4 {
			query := r.URL.Query()
			if _, ok := query["yaml"]; ok {
				config, _ = yaml.JSONToYAML(config)
				contentType = "text/x-yaml; charset=utf-8"
			} else if _, ok := query["pretty"]; ok {
				var raw json.RawMessage
				json.Unmarshal(config, &raw)
				config, _ = json.MarshalIndent(raw, "", "  ")
			}
		}
		header := rw.Header()
		header.Set("Content-Type", contentType)
		header.Set("Content-Length", strconv.FormatInt(int64(len(config)), 10))
		if http.MethodHead != r.Method {
			rw.Write(config)
		}
	case http.MethodPut:
		if c.readOnly {
			http.Error(rw, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
		body, err = yaml.JSONToYAML(body)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
		ctx := r.Context()
		pip, ok := pipe.GetPipeWithContext(ctx)
		if !ok {
			http.Error(rw, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
			return
		}

		err = pip.Reload(body)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
		rw.WriteHeader(http.StatusOK)
	}
}
