package log

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/wzshiming/pipe/configure/decode"
	"github.com/wzshiming/pipe/pipe/stdio/output"
)

const name = "log"

func init() {
	decode.Register(name, NewLogWithConfig)
}

type Config struct {
	Output  output.Output
	Handler http.Handler
}

func NewLogWithConfig(conf *Config) http.Handler {
	if conf.Output != nil {
		return handlers.CombinedLoggingHandler(conf.Output, conf.Handler)
	}
	return conf.Handler
}
