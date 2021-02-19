package file

import (
	"net/http"

	"github.com/pipeproxy/pipe/components/common/register"
)

const (
	name = "file"
)

func init() {
	register.Register(name, NewFileWithConfig)
}

type Config struct {
	AutoIndex bool     `json:",omitempty"`
	Indexes   []string `json:",omitempty"`
	Root      string
}

func NewFileWithConfig(conf *Config) http.Handler {
	return &FileServer{
		AutoIndex:  conf.AutoIndex,
		Indexes:    conf.Indexes,
		FileSystem: http.Dir(conf.Root),
	}
}
