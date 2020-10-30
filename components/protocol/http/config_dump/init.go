package config_dump

import (
	"net/http"

	"github.com/pipeproxy/pipe/components/common/register"
)

const (
	name = "config_dump"
)

func init() {
	register.Register(name, NewConfigDumpWithConfig)
}

type Config struct {
	ReadOnly bool
}

func NewConfigDumpWithConfig(conf *Config) http.Handler {
	return NewConfigDump(conf.ReadOnly)
}
