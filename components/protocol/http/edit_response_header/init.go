package edit_response_header

import (
	"net/http"

	"github.com/pipeproxy/pipe/components/common/register"
)

const (
	name = "edit_response_header"
)

func init() {
	register.Register(name, NewEditResponseHeaderWithConfig)
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

func NewEditResponseHeaderWithConfig(conf *Config) (http.Handler, error) {
	return NewEditResponseHeader(conf.Del, conf.Set, conf.Add)
}
