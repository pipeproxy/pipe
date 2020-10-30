package compress

import (
	"compress/gzip"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/pipeproxy/pipe/components/common/register"
)

const (
	name = "compress"
)

func init() {
	register.Register(name, NewLogWithConfig)
}

type Config struct {
	Level   int
	Handler http.Handler
}

func NewLogWithConfig(conf *Config) http.Handler {
	if conf.Level == 0 {
		conf.Level = gzip.DefaultCompression
	}
	return handlers.CompressHandlerLevel(conf.Handler, conf.Level)
}
