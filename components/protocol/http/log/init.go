package log

import (
	"net/http"

	"github.com/pipeproxy/pipe/components/common/register"
	"github.com/pipeproxy/pipe/components/stdio/output"
)

const (
	name = "log"
)

func init() {
	register.Register(name, NewLogWithConfig)
}

type Config struct {
	Output  output.Output
	Handler http.Handler
}

func NewLogWithConfig(conf *Config) http.Handler {
	return NewLog(conf.Handler)
}
