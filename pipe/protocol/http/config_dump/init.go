package config_dump

import (
	"net/http"

	"github.com/wzshiming/pipe/configure/decode"
)

const name = "config_dump"

func init() {
	decode.Register(name, NewConfigDumpWithConfig)
}

func NewConfigDumpWithConfig() http.Handler {
	return &ConfigDump{}
}
