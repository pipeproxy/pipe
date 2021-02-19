package file

import (
	"os"

	"github.com/pipeproxy/pipe/components/common/register"
	"github.com/pipeproxy/pipe/components/stdio/input"
)

const (
	name = "file"
)

func init() {
	register.Register(name, NewFileWithConfig)
}

type Config struct {
	Path string
}

func NewFileWithConfig(conf *Config) input.Input {
	return input.NewLazyReader(func() (input.Input, error) {
		f, err := os.Open(conf.Path)
		if err != nil {
			return nil, err
		}
		return input.NewReaderWithAutoClose(f, f), nil
	})
}
