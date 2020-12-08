package edit_request_header

import (
	"net/http"

	"github.com/pipeproxy/pipe/components/common/register"
)

const (
	name = "edit_request_header"
)

func init() {
	register.Register(name, NewEditRequestHeaderWithConfig)
}

type Pair struct {
	Key   string
	Value string
}

type Config struct {
	Del []string `json:",omitempty"`
	Set []Pair   `json:",omitempty"`
	Add []Pair   `json:",omitempty"`
}

func NewEditRequestHeaderWithConfig(conf *Config) (http.Handler, error) {
	return NewEditRequestHeader(conf.Del, conf.Set, conf.Add)
}
