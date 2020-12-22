package log

import (
	"github.com/pipeproxy/pipe/components/stream"

	"github.com/pipeproxy/pipe/components/common/register"
	"github.com/pipeproxy/pipe/components/stdio/output"
)

const (
	name = "log"
)

func init() {
	register.Register(name, NewLogWithConfig)
}

type Config struct {
	Output              output.Output
	Handler             stream.Handler
	OriginalDestination bool `json:",omitempty"`
}

func NewLogWithConfig(conf *Config) stream.Handler {
	return NewLog(conf.Handler, conf.Output, conf.OriginalDestination)
}
