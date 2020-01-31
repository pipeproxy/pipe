package config_dump

import (
	"net/http"

	"github.com/wzshiming/pipe/configure"
)

const name = "config_dump"

func init() {
	configure.Register(name, NewConfigDumpWithConfig)
}

func NewConfigDumpWithConfig() http.Handler {
	return &ConfigDump{}
}
