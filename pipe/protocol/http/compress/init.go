package compress

import (
	"compress/gzip"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/wzshiming/pipe/configure/decode"
)

const name = "compress"

func init() {
	decode.Register(name, NewLogWithConfig)
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
