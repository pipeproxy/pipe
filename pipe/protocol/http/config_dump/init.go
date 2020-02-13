package config_dump

import (
	"net/http"

	"github.com/wzshiming/pipe/configure/manager"
)

const name = "config_dump"

func init() {
	manager.Register(name, NewConfigDumpWithConfig)
}

func NewConfigDumpWithConfig() http.Handler {
	return &ConfigDump{}
}
