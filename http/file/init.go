package file

import (
	"net/http"

	"github.com/wzshiming/pipe/configure"
)

const name = "file"

func init() {
	configure.Register(name, NewFileWithConfig)
}

type Config struct {
	Root string
}

func NewFileWithConfig(conf *Config) http.Handler {
	return NewFile(conf.Root)
}
