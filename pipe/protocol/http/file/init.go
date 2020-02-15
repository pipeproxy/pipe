package file

import (
	"net/http"

	"github.com/wzshiming/pipe/configure/decode"
)

const name = "file"

func init() {
	decode.Register(name, NewFileWithConfig)
}

type Config struct {
	Root string
}

func NewFileWithConfig(conf *Config) http.Handler {
	return NewFile(conf.Root)
}
