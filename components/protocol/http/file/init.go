package file

import (
	"net/http"

	"github.com/wzshiming/pipe/components/common/register"
)

const name = "file"

func init() {
	register.Register(name, NewFileWithConfig)
}

type Config struct {
	Root string
}

func NewFileWithConfig(conf *Config) http.Handler {
	return NewFile(conf.Root)
}
