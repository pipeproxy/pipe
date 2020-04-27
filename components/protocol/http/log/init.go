package log

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/components/stdio/output"
)

const name = "log"

func init() {
	register.Register(name, NewLogWithConfig)
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