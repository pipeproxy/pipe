package file

import (
	"net/http"

	"github.com/wzshiming/pipe/configure/manager"
)

const name = "file"

func init() {
	manager.Register(name, NewFileWithConfig)
}

type Config struct {
	Root string
}

func NewFileWithConfig(conf *Config) http.Handler {
	return NewFile(conf.Root)
}
