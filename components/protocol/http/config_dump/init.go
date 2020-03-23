package config_dump

import (
	"net/http"

	"github.com/wzshiming/pipe/components/common/register"
)

const name = "config_dump"

func init() {
	register.Register(name, NewConfigDumpWithConfig)
}

func NewConfigDumpWithConfig() http.Handler {
	return &ConfigDump{}
}
